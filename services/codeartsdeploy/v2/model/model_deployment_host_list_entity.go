package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeploymentHostListEntity struct {

	// 主机id列表
	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o DeploymentHostListEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentHostListEntity struct{}"
	}

	return strings.Join([]string{"DeploymentHostListEntity", string(data)}, " ")
}
