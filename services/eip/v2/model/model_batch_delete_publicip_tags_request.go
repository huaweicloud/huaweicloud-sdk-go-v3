/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Request Object
type BatchDeletePublicipTagsRequest struct {
	PublicipId string `json:"publicip_id"`
	Body *BatchDeletePublicipTagsRequestBody `json:"body,omitempty"`
}
