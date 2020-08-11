/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Request Object
type BatchCreatePublicipTagsRequest struct {
	PublicipId string `json:"publicip_id"`
	Body *BatchCreatePublicipTagsRequestBody `json:"body,omitempty"`
}
