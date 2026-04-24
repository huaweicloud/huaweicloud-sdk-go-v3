package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscriptionDataType 订阅的数据类型是否包含DDL/DML语句，返回值： true：是。 false：否。
type SubscriptionDataType struct {

	// 数据操作语言，取值： true：订阅DML false：不订阅DML
	IsDmlSubscribed bool `json:"is_dml_subscribed"`

	// 数据定义语言，取值： true：订阅DDL false：不订阅DDL
	IsDdlSubscribed bool `json:"is_ddl_subscribed"`
}

func (o SubscriptionDataType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionDataType struct{}"
	}

	return strings.Join([]string{"SubscriptionDataType", string(data)}, " ")
}
