package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkPolicyDocument 中心网络策略文档。
type CentralNetworkPolicyDocument struct {

	// 中心网络默认平面的名字。
	DefaultPlane string `json:"default_plane"`

	// 中心网络平面列表。
	Planes []CentralNetworkPlaneDocument `json:"planes"`

	// 中心网络ER实例列表。
	ErInstances *[]AssociateErInstanceDocument `json:"er_instances,omitempty"`
}

func (o CentralNetworkPolicyDocument) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkPolicyDocument struct{}"
	}

	return strings.Join([]string{"CentralNetworkPolicyDocument", string(data)}, " ")
}
