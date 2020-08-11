/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model

// Request Object
type CreateTranscodingsTemplateRequest struct {
	Body *StreamTranscodingTemplate `json:"body,omitempty"`
}
