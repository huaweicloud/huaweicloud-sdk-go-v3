package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Association struct {

	// 网卡绑定的eipId
	PublicIpAddress *string `json:"public_ip_address,omitempty"`
}

func (o Association) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Association struct{}"
	}

	return strings.Join([]string{"Association", string(data)}, " ")
}
