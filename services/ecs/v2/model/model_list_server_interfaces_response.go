/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Response Object
type ListServerInterfacesResponse struct {
	// 云服务器网卡信息列表
	InterfaceAttachments []InterfaceAttachment `json:"interfaceAttachments,omitempty"`
}
