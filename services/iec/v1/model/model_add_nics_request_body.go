package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 添加网卡请求体
type AddNicsRequestBody struct {
	// 虚拟私有云ID。

	VpcId string `json:"vpc_id"`
	// 安全组ID列表。

	SecurityGroups []BaseId `json:"security_groups"`
	// 子网ID。  当subnet_id提供时，则在该子网下创建nic_num个网卡； 不输入，则自动分配subnet。 当添加网卡的VPC为手动规划VPC时，subnet_id必填。

	SubnetId *string `json:"subnet_id,omitempty"`
	// 待添加网卡数量。

	NicNum int32 `json:"nic_num"`
}

func (o AddNicsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddNicsRequestBody struct{}"
	}

	return strings.Join([]string{"AddNicsRequestBody", string(data)}, " ")
}
