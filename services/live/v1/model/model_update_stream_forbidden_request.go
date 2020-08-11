/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model

// Request Object
type UpdateStreamForbiddenRequest struct {
	SpecifyProject string `json:"specify_project,omitempty"`
	Body *StreamForbiddenSetting `json:"body,omitempty"`
}
