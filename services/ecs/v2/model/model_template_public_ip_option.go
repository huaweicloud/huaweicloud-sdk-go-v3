package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplatePublicIpOption struct {

	// 弹性公网IP类型
	PublicipType *string `json:"publicip_type,omitempty"`

	// 弹性公网IP计费类型
	ChargingMode *string `json:"charging_mode,omitempty"`

	Bandwidth *TemplateBandwidthOption `json:"bandwidth,omitempty"`

	// EIP是否随实例一同释放
	DeleteOnTermination *bool `json:"delete_on_termination,omitempty"`
}

func (o TemplatePublicIpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplatePublicIpOption struct{}"
	}

	return strings.Join([]string{"TemplatePublicIpOption", string(data)}, " ")
}
