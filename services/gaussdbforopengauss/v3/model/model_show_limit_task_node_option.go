package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowLimitTaskNodeOption struct {

	// 节点id。
	NodeId string `json:"node_id"`

	// 该节点执行的sql语句id。
	SqlId string `json:"sql_id"`
}

func (o ShowLimitTaskNodeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLimitTaskNodeOption struct{}"
	}

	return strings.Join([]string{"ShowLimitTaskNodeOption", string(data)}, " ")
}
