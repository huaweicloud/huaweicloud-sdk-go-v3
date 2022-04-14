package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnfreezeSubCustomersReq struct {
	// 需要解冻的客户账号ID列表。 您可以调用查询客户列表接口获取customer_id。

	CustomerIds []string `json:"customer_ids"`
	// 解冻原因。

	Reason string `json:"reason"`
}

func (o UnfreezeSubCustomersReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnfreezeSubCustomersReq struct{}"
	}

	return strings.Join([]string{"UnfreezeSubCustomersReq", string(data)}, " ")
}
