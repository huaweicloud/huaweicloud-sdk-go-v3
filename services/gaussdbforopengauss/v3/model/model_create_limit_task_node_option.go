package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateLimitTaskNodeOption struct {

	// 节点id。
	NodeId string `json:"node_id"`

	// 该节点执行的sql语句id。如果类型为SQL_ID，必须与limit_type_value值一致。
	SqlId string `json:"sql_id"`
}

func (o CreateLimitTaskNodeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLimitTaskNodeOption struct{}"
	}

	return strings.Join([]string{"CreateLimitTaskNodeOption", string(data)}, " ")
}
