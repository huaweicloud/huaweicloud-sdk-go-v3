package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowResourcesDetailResponseBody struct {

	// 配额资源类型，当前配额类型仅支持实例类型（instance）。
	Type string `json:"type"`

	// 实例类型： - 若值为“Single”，则表示单节点实例配额信息。 - 若值为“ReplicaSet”，则表示副本集实例配额信息。 - 若值为“Sharding”，则表示集群实例配额信息。
	Mode string `json:"mode"`

	// 当前配额值。
	Quota int32 `json:"quota"`

	// 已使用的资源数。
	Used int32 `json:"used"`
}

func (o ShowResourcesDetailResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourcesDetailResponseBody struct{}"
	}

	return strings.Join([]string{"ShowResourcesDetailResponseBody", string(data)}, " ")
}
