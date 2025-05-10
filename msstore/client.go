package msstore

import (
	"crypto/tls"
	"encoding/json"
	"fmt"

	"ms_store_api/msstore/internal/genurl"

	"github.com/labstack/gommon/log"
	"resty.dev/v3"
)

// FetchProductDetails 从Microsoft Store API获取产品信息
// 参数:
//   - id: 要获取详情的产品ID
//
// 返回值:
//   - string: UWP应用的产品文件名
//   - error: 过程中遇到的任何错误
//
// 使用示例:
//
//	productInfo, err := FetchProductDetails("9WZDNCRFHVN5") // 获取WhatsApp的详情
//	if err != nil {
//	    log.Fatal(err)
//	}
func FetchProductDetails(id string) (string, error) {
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	api_url := fmt.Sprintf("https://storeedgefd.dsx.mp.microsoft.com/v9.0/products/%s", id)

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"market":       "US",
			"locale":       "en-us",
			"deviceFamily": "Windows.Desktop",
		}).
		SetHeader("Accept", "application/json").
		Get(api_url)

	if err != nil {
		log.Error(err)
	}

	var res ResponseItem
	if err := json.Unmarshal(resp.Bytes(), &res); err != nil {
		return "", fmt.Errorf("Failed to parse response: %v", err)
	}

	if res.Payload == nil {
		return "", fmt.Errorf("Invalid product ID")
	}

	if len(res.Payload.Skus) > 0 && res.Payload.Skus[0].FulfillmentData != "" {
		// UWP应用
		res, err := genurl.GenUWPUrl(client, res.Payload.Skus[0].FulfillmentData)
		return res, err
	}

	// 非UWP应用
	return "", nil
}
