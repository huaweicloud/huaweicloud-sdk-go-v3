package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VpcId 创建网卡所属的 VPC ID，可通过 VPC API 查询：https://support.huaweicloud.com/api-vpc/vpc_api01_0003.html。
type VpcId struct {
}

func (o VpcId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcId struct{}"
	}

	return strings.Join([]string{"VpcId", string(data)}, " ")
}
