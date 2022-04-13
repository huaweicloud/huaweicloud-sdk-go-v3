package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FollowerMigrateRequest struct {
	// 备机节点Id

	NodeId string `json:"nodeId"`
	// 要迁入的可用区code

	AzCode string `json:"azCode"`
}

func (o FollowerMigrateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FollowerMigrateRequest struct{}"
	}

	return strings.Join([]string{"FollowerMigrateRequest", string(data)}, " ")
}
