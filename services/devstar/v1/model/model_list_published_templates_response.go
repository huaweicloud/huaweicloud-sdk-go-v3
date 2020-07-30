/*
    * Devstar
    *
    * Devstar API
    *
*/

package model

// Response Object
type ListPublishedTemplatesResponse struct {
	// 返回模板的数量
	Count int32 `json:"count,omitempty"`
	// 返回模板的列表
	Templates []TemplateSimpleInfo `json:"templates,omitempty"`
}
