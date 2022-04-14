package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 路由表关联子网请求体
type AssociateSubnetRequestBody struct {
	// 子网ID

	SubnetIds []string `json:"subnet_ids"`
}

func (o AssociateSubnetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateSubnetRequestBody struct{}"
	}

	return strings.Join([]string{"AssociateSubnetRequestBody", string(data)}, " ")
}
