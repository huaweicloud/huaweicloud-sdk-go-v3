/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Response Object
type ListServerBlockDevicesResponse struct {
	AttachableQuantity *BlockDeviceAttachableQuantity `json:"attachableQuantity,omitempty"`
	// 云服务器挂载信息列表。
	VolumeAttachments []ServerBlockDevice `json:"volumeAttachments,omitempty"`
}
