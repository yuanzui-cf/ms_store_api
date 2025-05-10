package msstore

// ResponseItem 表示从Microsoft Store API的顶层响应
type ResponseItem struct {
	Payload *ProductDetails `json:"Payload,omitempty"`
}

// ProductDetails 包含Microsoft Store中产品的信息
type ProductDetails struct {
	Skus []Sku `json:"Skus,omitempty"`
}

// Sku 表示具有履行数据的产品特定SKU
type Sku struct {
	FulfillmentData string `json:"FulfillmentData,omitempty"`
}
