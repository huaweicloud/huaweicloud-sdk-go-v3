package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyBindEipRequest struct {

	// 待绑定的弹性公网IP地址。
	PublicIp string `json:"public_ip"`

	// 弹性公网IP地址对应的ID。
	PublicIpId string `json:"public_ip_id"`
}

func (o ModifyBindEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyBindEipRequest struct{}"
	}

	return strings.Join([]string{"ModifyBindEipRequest", string(data)}, " ")
}
