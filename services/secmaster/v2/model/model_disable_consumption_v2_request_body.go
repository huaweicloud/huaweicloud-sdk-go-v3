package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DisableConsumptionV2RequestBody struct {

	// 表格id
	TableId string `json:"table_id"`
}

func (o DisableConsumptionV2RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableConsumptionV2RequestBody struct{}"
	}

	return strings.Join([]string{"DisableConsumptionV2RequestBody", string(data)}, " ")
}
