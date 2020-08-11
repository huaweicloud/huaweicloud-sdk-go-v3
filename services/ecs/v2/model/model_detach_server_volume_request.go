/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type DetachServerVolumeRequest struct {
	ServerId string `json:"server_id"`
	VolumeId string `json:"volume_id"`
	DeleteFlag string `json:"delete_flag,omitempty"`
}
