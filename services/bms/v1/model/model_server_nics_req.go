package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ServerNicsReq struct {

	// 网卡的子网ID
	SubnetId string `json:"subnet_id"`

	//
	IpAddress *string `json:"ip_address,omitempty"`

	SecurityGroups *SecurityGroupInfo `json:"security_groups,omitempty"`
}

func (o ServerNicsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerNicsReq struct{}"
	}

	return strings.Join([]string{"ServerNicsReq", string(data)}, " ")
}
