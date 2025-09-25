package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTransactionRequestBodyTransactionQueryOption **参数解释**: 查询事务筛选条件。 **约束限制**: 不涉及。
type ListTransactionRequestBodyTransactionQueryOption struct {

	// **参数解释**: 事务执行时长，单位：秒。 **约束限制**: 不涉及。 **取值范围**: 非负整数。 **默认取值**: 0
	ExecTime *string `json:"exec_time,omitempty"`

	// **参数解释**: 事务xlog日志大小：单位byte。 **约束限制**: 不涉及。 **取值范围**: 非负整数。 **默认取值**: 0
	XlogQuantity *string `json:"xlog_quantity,omitempty"`

	// **参数解释**: 数据库名。 **约束限制**: 不涉及。
	Datnames *[]string `json:"datnames,omitempty"`

	// **参数解释**: 用户名。 **约束限制**: 不涉及。
	Usenames *[]string `json:"usenames,omitempty"`

	// **参数解释**: 用户发起事务请求的客户端地址。 **约束限制**: 不涉及。
	ClientAddrs *[]string `json:"client_addrs,omitempty"`
}

func (o ListTransactionRequestBodyTransactionQueryOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransactionRequestBodyTransactionQueryOption struct{}"
	}

	return strings.Join([]string{"ListTransactionRequestBodyTransactionQueryOption", string(data)}, " ")
}
