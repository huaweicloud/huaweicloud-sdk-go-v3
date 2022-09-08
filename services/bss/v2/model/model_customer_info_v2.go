package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerInfoV2 struct {

	// 客户账号ID。您可以调用[查询客户列表](https://support.huaweicloud.com/api-bpconsole/mc_00021.html)接口获取customer_id。
	CustomerId string `json:"customer_id"`
}

func (o CustomerInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerInfoV2 struct{}"
	}

	return strings.Join([]string{"CustomerInfoV2", string(data)}, " ")
}
