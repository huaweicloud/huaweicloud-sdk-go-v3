package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BindEipOpenRequest struct {

	// 公共ip。
	PublicIp *string `json:"public_ip,omitempty"`

	// 公共ip id。
	PublicIpId *string `json:"public_ip_id,omitempty"`
}

func (o BindEipOpenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindEipOpenRequest struct{}"
	}

	return strings.Join([]string{"BindEipOpenRequest", string(data)}, " ")
}
