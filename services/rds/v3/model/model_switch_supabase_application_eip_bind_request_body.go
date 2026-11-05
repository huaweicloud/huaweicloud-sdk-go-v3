package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchSupabaseApplicationEipBindRequestBody EIP绑解请求体。
type SwitchSupabaseApplicationEipBindRequestBody struct {

	// true绑定，false解绑。is_bind=true时public_ip_id和public_ip_address必填。
	IsBind bool `json:"is_bind"`

	// 弹性公网IP地址。
	PublicIpAddress *string `json:"public_ip_address,omitempty"`

	// 弹性公网IP ID（UUID）。
	PublicIpId *string `json:"public_ip_id,omitempty"`
}

func (o SwitchSupabaseApplicationEipBindRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSupabaseApplicationEipBindRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchSupabaseApplicationEipBindRequestBody", string(data)}, " ")
}
