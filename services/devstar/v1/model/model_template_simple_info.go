/*
    * Devstar
    *
    * Devstar API
    *
*/

package model

type TemplateSimpleInfo struct {
	// 模板id
	Id string `json:"id,omitempty"`
	// 模板名
	Title string `json:"title,omitempty"`
	// 模板描述
	Description string `json:"description,omitempty"`
}
