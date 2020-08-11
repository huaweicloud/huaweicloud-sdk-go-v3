/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// Request Object
type ListBandwidthsRequest struct {
	Marker string `json:"marker,omitempty"`
	Limit int32 `json:"limit,omitempty"`
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	ShareType string `json:"share_type,omitempty"`
}
