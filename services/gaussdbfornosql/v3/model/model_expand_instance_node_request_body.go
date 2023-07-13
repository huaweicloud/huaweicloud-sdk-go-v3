package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExpandInstanceNodeRequestBody struct {

	// 新增加的节点数量。
	Num int32 `json:"num"`

	// 扩容的节点所使用的子网的ID。 - 该参数仅只支持GaussDB(for Cassandra)数据库实例扩容节点时传入。 - 所传入的子网ID必须属于实例当前所在的VPC。 - 不传该参数时，系统会在当前实例所使用的子网中为当前扩容的节点选择一个IP容量较为充足的子网。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 创建包周期实例时可指定，表示是否自动从账户中支付，此字段不影响自动续订的支付方式。 - true，表示自动从账户中支付。 - false，表示手动从账户中支付，默认为该方式。
	IsAutoPay *string `json:"is_auto_pay,omitempty"`
}

func (o ExpandInstanceNodeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandInstanceNodeRequestBody struct{}"
	}

	return strings.Join([]string{"ExpandInstanceNodeRequestBody", string(data)}, " ")
}
