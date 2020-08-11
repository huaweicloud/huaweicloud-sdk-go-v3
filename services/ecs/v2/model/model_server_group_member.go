/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 云服务器组添加、删除成员列表
type ServerGroupMember struct {
	// 云服务器UUID。
	InstanceUuid string `json:"instance_uuid"`
}
