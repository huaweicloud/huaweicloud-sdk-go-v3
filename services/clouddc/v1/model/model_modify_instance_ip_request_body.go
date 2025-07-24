package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyInstanceIpRequestBody 修改服务器ip请求体。
type ModifyInstanceIpRequestBody struct {

	// 创建网卡所属的 VPC ID，可通过 VPC API 查询：https://support.huaweicloud.com/api-vpc/vpc_api01_0003.html。
	VpcId *string `json:"vpc_id,omitempty"`

	// 网卡信息
	Nics *[]ModifyInstanceIpRequestBodyNics `json:"nics,omitempty"`
}

func (o ModifyInstanceIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyInstanceIpRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyInstanceIpRequestBody", string(data)}, " ")
}
