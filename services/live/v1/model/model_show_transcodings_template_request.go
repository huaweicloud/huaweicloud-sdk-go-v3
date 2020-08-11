/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model

// Request Object
type ShowTranscodingsTemplateRequest struct {
	Domain string `json:"domain"`
	AppName string `json:"app_name,omitempty"`
	Page int32 `json:"page,omitempty"`
	Size int32 `json:"size,omitempty"`
}
