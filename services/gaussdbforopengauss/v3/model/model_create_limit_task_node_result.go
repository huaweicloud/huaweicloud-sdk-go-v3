package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateLimitTaskNodeResult struct {

	// 节点id。
	NodeId *string `json:"node_id,omitempty"`

	// 该节点执行的sql语句id。
	SqlId *string `json:"sql_id,omitempty"`
}

func (o CreateLimitTaskNodeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLimitTaskNodeResult struct{}"
	}

	return strings.Join([]string{"CreateLimitTaskNodeResult", string(data)}, " ")
}
