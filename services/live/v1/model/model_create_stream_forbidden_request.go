/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model

// Request Object
type CreateStreamForbiddenRequest struct {
	SpecifyProject string `json:"specify_project,omitempty"`
	Body *StreamForbiddenSetting `json:"body,omitempty"`
}
