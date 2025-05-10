package xml

import (
	"fmt"

	"github.com/beevik/etree"
)

// ExtractFilenames 从XML响应中提取文件名信息
// 参数:
//   - files: XML中的文件元素数组
//
// 返回:
//   - map[string]FileInfo: 以节点ID为键，FileInfo为值的映射
//   - error: 处理过程中的错误
//
// 用法示例:
// ```
// doc := etree.NewDocument()
// err := doc.ReadFromString(xmlString)
//
//	if err != nil {
//	    return err
//	}
//
// files := doc.FindElements("//File")
// filenamesMap, err := ExtractFilenames(files)
//
//	if err != nil {
//	    return err
//	}
//
// ```
func ExtractFilenames(files []*etree.Element) (map[string]FileInfo, error) {
	filenamesMap := make(map[string]FileInfo)

	fmt.Println(files)
	// 遍历每个文件元素
	for _, file := range files {
		// 检查父节点是否存在
		parent := file.Parent()
		if parent == nil {
			continue
		}
		fmt.Println(1)

		// 检查祖父节点是否存在
		grandParent := parent.Parent()
		if grandParent == nil {
			continue
		}
		fmt.Println(2)

		// 从祖父节点中查找ID元素
		idNode := grandParent.FindElement(".//ID")
		if idNode == nil || idNode.Text() == "" {
			continue
		}
		fmt.Println(3)

		nodeID := idNode.Text()

		// 如果当前元素有子元素，使用第一个子元素
		if len(file.ChildElements()) > 0 {
			file = file.ChildElements()[0]
		}

		// 提取所需属性
		installerSpecificIdentifier := file.SelectAttr("InstallerSpecificIdentifier")
		fileName := file.SelectAttr("FileName")
		modified := file.SelectAttr("Modified")

		// 确保所有需要的属性都存在
		if installerSpecificIdentifier != nil && fileName != nil && modified != nil {
			filenamesMap[nodeID] = FileInfo{
				Identifier: fmt.Sprintf("%s_%s", installerSpecificIdentifier.Value, fileName.Value),
				Modified:   modified.Value,
			}
		}
	}

	// 检查是否成功提取了任何文件名
	if len(filenamesMap) == 0 {
		return nil, fmt.Errorf("无法从响应中提取文件名")
	}

	return filenamesMap, nil
}

// ExtractUpdateIdentities 从XML中提取更新ID和修订号
// 参数:
//   - fragmentNodes: XML中的片段节点元素数组
//   - filenamesMap: 从ExtractFilenames获取的文件信息映射
//
// 返回:
//   - map[string]UpdateIdentity: 以文件名为键，更新标识为值的映射
//   - map[string]string: 以文件名为键，修改时间为值的映射
//   - error: 处理过程中的错误
//
// 用法示例:
// ```
// doc := etree.NewDocument()
// err := doc.ReadFromString(xmlString)
//
//	if err != nil {
//	    return err
//	}
//
// files := doc.FindElements("//File")
// filenamesMap, err := ExtractFilenames(files)
//
//	if err != nil {
//	    return err
//	}
//
// fragmentNodes := doc.FindElements("//Fragment")
// identities, nameModified, err := ExtractUpdateIdentities(fragmentNodes, filenamesMap)
//
//	if err != nil {
//	    return err
//	}
//
// ```
func ExtractUpdateIdentities(fragmentNodes []*etree.Element, filenamesMap map[string]FileInfo) (map[string]UpdateIdentity, map[string]string, error) {
	identities := make(map[string]UpdateIdentity)
	nameModified := make(map[string]string)

	// 遍历每个片段节点
	for _, fragmentNode := range fragmentNodes {
		// 获取父节点并检查其是否存在
		parentNode := fragmentNode.Parent()
		if parentNode == nil {
			continue
		}

		// 获取祖父节点并检查其是否存在
		grandParentNode := parentNode.Parent()
		if grandParentNode == nil {
			continue
		}

		// 获取曾祖父节点并检查其是否存在
		greatGrandParentNode := grandParentNode.Parent()
		if greatGrandParentNode == nil {
			continue
		}

		// 从曾祖父节点中查找ID元素
		idNode := greatGrandParentNode.FindElement(".//ID")
		if idNode == nil || idNode.Text() == "" {
			continue
		}

		fnID := idNode.Text()

		// 从filenamesMap中获取对应的文件信息
		fileInfo, exists := filenamesMap[fnID]
		if !exists {
			continue
		}

		fileName := fileInfo.Identifier
		modified := fileInfo.Modified

		// 从祖父节点中查找UpdateIdentity元素
		topNode := grandParentNode.FindElement(".//UpdateIdentity")
		if topNode == nil {
			continue
		}

		// 提取更新ID和修订号属性
		updateID := topNode.SelectAttr("UpdateID")
		revNum := topNode.SelectAttr("RevisionNumber")

		// 确保所有需要的属性都存在
		if updateID != nil && revNum != nil {
			nameModified[fileName] = modified
			identities[fileName] = UpdateIdentity{
				UpdateID:       updateID.Value,
				RevisionNumber: revNum.Value,
			}
		}
	}

	return identities, nameModified, nil
}
