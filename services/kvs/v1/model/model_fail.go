package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Fail 失败的操作指示。
type Fail struct {

	// 失败的操作标识，1个或多个。
	OperId int32 `bson:"oper_id"`

	// 处理失败操作提示。
	Status string `bson:"status"`
}

func (o Fail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Fail struct{}"
	}

	return strings.Join([]string{"Fail", string(data)}, " ")
}
