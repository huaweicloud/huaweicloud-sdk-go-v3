package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateEcnWithVpcRequestBody struct {

	// 本端子网列表
	LocalSubnetList *[]string `json:"local_subnet_list,omitempty"`

	// 是否刷新对端子网路由
	RefreshRemoteSubnetRoute bool `json:"refresh_remote_subnet_route"`
}

func (o UpdateEcnWithVpcRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEcnWithVpcRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEcnWithVpcRequestBody", string(data)}, " ")
}
