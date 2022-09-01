package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlErrorLogList struct {

	// 节点ID。
	NodeId *string `json:"node_id,omitempty" xml:"node_id"`

	// 日期时间UTC时间。
	Time *string `json:"time,omitempty" xml:"time"`

	// 日志级别。
	Level *string `json:"level,omitempty" xml:"level"`

	// 错误日志内容。
	Content *string `json:"content,omitempty" xml:"content"`
}

func (o MysqlErrorLogList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlErrorLogList struct{}"
	}

	return strings.Join([]string{"MysqlErrorLogList", string(data)}, " ")
}
