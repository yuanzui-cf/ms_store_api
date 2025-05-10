package genurl

// FulfillmentDataStruct 包含下载 UWP 应用所需的标识符
type FulfillmentDataStruct struct {
	WuCategoryId      string `json:"WuCategoryId,omitempty"`      // Windows Update 类别 ID
	PackageFamilyName string `json:"PackageFamilyName,omitempty"` // UWP 应用家族名称
}
