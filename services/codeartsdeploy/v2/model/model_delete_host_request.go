package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHostRequest Request Object
type DeleteHostRequest struct {

	// 主机集群id
	GroupId string `json:"group_id"`

	// 主机id
	HostId string `json:"host_id"`
}

func (o DeleteHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostRequest struct{}"
	}

	return strings.Join([]string{"DeleteHostRequest", string(data)}, " ")
}
