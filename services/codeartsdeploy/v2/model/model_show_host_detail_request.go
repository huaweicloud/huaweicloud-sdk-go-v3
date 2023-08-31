package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHostDetailRequest Request Object
type ShowHostDetailRequest struct {

	// 主机集群id
	GroupId string `json:"group_id"`

	// 主机id
	HostId string `json:"host_id"`
}

func (o ShowHostDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHostDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowHostDetailRequest", string(data)}, " ")
}
