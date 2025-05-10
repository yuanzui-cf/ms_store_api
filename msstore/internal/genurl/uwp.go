package genurl

import (
	"encoding/json"
	"fmt"
	"html"
	"regexp"
	"strings"

	"ms_store_api/msstore/internal/xml"

	"github.com/beevik/etree"
	"resty.dev/v3"
)

// GenUWPUrl 通过与Microsoft的Windows Update服务交互生成用于下载UWP应用程序的URL
// 参数:
//   - client: resty客户端，用于HTTP请求
//   - data: 来自产品详情的fulfillment数据字符串
//
// 返回:
//   - string: UWP应用程序的文件名
//   - error: 处理过程中遇到的任何错误
//
// 用法:
//
//	client := resty.New()
//	filename, err := GenUWPUrl(client, fulfillmentData)
func GenUWPUrl(client *resty.Client, data string) (string, error) {
	var list FulfillmentDataStruct
	if err := json.Unmarshal([]byte(data), &list); err != nil {
		return "", fmt.Errorf("Failed to parse fulfillment data: %v", err)
	}

	categoryId, fileName := list.WuCategoryId, strings.Split(list.PackageFamilyName, "_")[0]
	releaseType := "retail"

	// 1. 从Microsoft服务获取加密Cookie
	var res_text string

	if cookieTemplate, err := DataDir.ReadFile("data/GetCookie.xml"); err != nil {
		return "", fmt.Errorf("Failed to read cookie template: %v", err)
	} else {
		resp, err := client.R().
			SetHeader("Content-Type", "application/soap+xml; charset=utf-8").
			SetBody(cookieTemplate).
			Post("https://fe3cr.delivery.mp.microsoft.com/ClientWebService/client.asmx")
		if err != nil {
			return "", err
		}

		res_text = resp.String()
	}

	doc := etree.NewDocument()
	if err := doc.ReadFromString(res_text); err != nil {
		return "", fmt.Errorf("Failed to parse response: %v", err)
	}

	encryptedCookie := doc.FindElement("//EncryptedData")
	if encryptedCookie == nil {
		return "", fmt.Errorf("Failed to find encrypted cookie")
	}

	cookie := encryptedCookie.Text()

	// 2. 使用Cookie和分类ID请求IDs和文件名
	if wuidTemplete, err := DataDir.ReadFile("data/WUIDRequest.xml"); err != nil {
		return "", fmt.Errorf("Failed to read WUID request template: %v", err)
	} else {
		wuidTemplete := string(wuidTemplete)
		wuidTemplete = strings.ReplaceAll(wuidTemplete, "{0}", cookie)
		wuidTemplete = strings.ReplaceAll(wuidTemplete, "{1}", categoryId)
		wuidTemplete = strings.ReplaceAll(wuidTemplete, "{2}", releaseType)

		resp, err := client.R().
			SetHeader("Content-Type", "application/soap+xml; charset=utf-8").
			SetBody(wuidTemplete).
			Post("https://fe3cr.delivery.mp.microsoft.com/ClientWebService/client.asmx")
		if err != nil {
			return "", err
		}

		res_text = resp.String()
	}

	doc = etree.NewDocument()
	if err := doc.ReadFromString(html.UnescapeString(res_text)); err != nil {
		return "", fmt.Errorf("Failed to parse response: %v", err)
	}

	files := doc.FindElements("//Files")

	// 3. 从响应中收集文件名
	filenamesMap, err := xml.ExtractFilenames(files)
	if err != nil {
		return "", err
	}

	if len(filenamesMap) == 0 {
		return "", fmt.Errorf("Server returned an empty list")
	}

	// 4. 从SecuredFragment元素解析更新IDs
	// 查找所有包含SecuredFragment的节点
	fragmentNodes := doc.FindElements("//SecuredFragment")

	identitiesMap, nameModifiedMap, err := xml.ExtractUpdateIdentities(fragmentNodes, filenamesMap)
	if err != nil {
		return "", err
	}

	// 5. 在可用选项中选择最佳文件

	fmt.Println(fileName)
	fmt.Println(filenamesMap)
	fmt.Println(identitiesMap)
	fmt.Println(nameModifiedMap)

	return "", nil
}

// ParseDict 解析字典数据并返回最佳匹配的文件名
//
// 参数:
//   - dict: 键为文件名，值为修改信息的映射
//   - fileName: 要匹配的文件名基础
//
// 返回:
//   - string: 最佳匹配的文件名
//   - error: 处理过程中的错误
func ParseDict(dict map[string]string, fileName string) (string, error) {
	// 准备文件名进行匹配
	base := CleanName(strings.Split(fileName, "-")[0])

	// 创建编译正则表达式，排除BlockMap文件
	re, err := regexp.Compile(`.+\.BlockMap`)
	if err != nil {
		return "", fmt.Errorf("编译正则表达式失败: %v", err)
	}

	// 构建用于查找的结构化数据
	files := make(map[string]string)
	for key, _ := range dict {
		if !re.MatchString(key) {
			parts := strings.Split(key, "_")
			if len(parts) >= 3 {
				nameBase := CleanName(parts[0])
				arch := parts[2]
				extParts := strings.Split(parts[len(parts)-1], ".")

				if len(extParts) >= 2 {
					ext := strings.ToLower(extParts[1])
					// 优先选择appx和x64架构
					if nameBase == base && (ext == "appx" || ext == "appxbundle") && (arch == "x64" || arch == "neutral") {
						return key, nil
					}

					// 保存所有可能的匹配项
					if nameBase == base {
						files[key] = ext + "_" + arch
					}
				}
			}
		}
	}

	// 如果找到匹配项，返回第一个
	if len(files) > 0 {
		for key := range files {
			return key, nil
		}
	}

	return "", fmt.Errorf("未找到匹配的文件")
}
