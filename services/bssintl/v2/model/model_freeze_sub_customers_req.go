package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FreezeSubCustomersReq struct {

	// 需要冻结的客户账号ID列表。 您可以调用查询客户列表接口获取customer_id。
	CustomerIds []string `json:"customer_ids"`

	// 冻结原因。
	Reason string `json:"reason"`
}

func (o FreezeSubCustomersReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FreezeSubCustomersReq struct{}"
	}

	return strings.Join([]string{"FreezeSubCustomersReq", string(data)}, " ")
}
