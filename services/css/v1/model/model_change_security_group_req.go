package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeSecurityGroupReq struct {

	// 期望安全组的ID。如何获取安全组ID，请参见[[《查询安全组列表》](https://support.huaweicloud.com/api-vpc/vpc_sg01_0003.html)](tag:hc)[[《查询安全组列表》](https://support.huaweicloud.com/intl/zh-cn/api-vpc/vpc_sg01_0003.html)](tag:hk,hws_hk)
	SecurityGroupIds string `json:"security_group_ids"`
}

func (o ChangeSecurityGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSecurityGroupReq struct{}"
	}

	return strings.Join([]string{"ChangeSecurityGroupReq", string(data)}, " ")
}
