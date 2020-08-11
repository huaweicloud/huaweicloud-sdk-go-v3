/*
    * LiveAPI
    *
    * 直播服务源站所有接口
    *
*/

package model

// Request Object
type CreateRecordConfigRequest struct {
	Body *RecordConfigInfo `json:"body,omitempty"`
}
