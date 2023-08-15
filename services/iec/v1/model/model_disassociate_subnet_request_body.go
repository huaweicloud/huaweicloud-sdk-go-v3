package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateSubnetRequestBody 路由表解关联子网请求体
type DisassociateSubnetRequestBody struct {

	// 子网ID
	SubnetIds []string `json:"subnet_ids"`
}

func (o DisassociateSubnetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateSubnetRequestBody struct{}"
	}

	return strings.Join([]string{"DisassociateSubnetRequestBody", string(data)}, " ")
}
