/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model

// Request Object
type DeleteStreamForbiddenRequest struct {
	SpecifyProject string `json:"specify_project,omitempty"`
	Domain string `json:"domain"`
	AppName string `json:"app_name"`
	StreamName string `json:"stream_name"`
}
