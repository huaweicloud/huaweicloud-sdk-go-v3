package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePublicIpRequestBody struct {

	// 公网访问绑定的ELB的EIP的ID，当Redis版本为3.0时，该参数为必填参数。
	PublicipId *string `json:"publicip_id,omitempty"`

	// 是否开启SSL，仅在开启SSL时有值，当Redis版本为3.0时，该参数为必填参数。
	EnableSsl *bool `json:"enable_ssl,omitempty"`

	// 公网访问绑定的ELB ID，当Redis版本为4.0，5.0，6.0和企业版时，该参数为必填参数。
	ElbId *string `json:"elb_id,omitempty"`
}

func (o UpdatePublicIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicIpRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePublicIpRequestBody", string(data)}, " ")
}
