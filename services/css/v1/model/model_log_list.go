package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 日志列表。
type LogList struct {

	// 日志内容。
	Content *string `json:"content,omitempty"`

	// 当条日志记录时间，格式为ISO8601：CCYY-MM-DDThh:mm:ss。
	Date *string `json:"date,omitempty"`

	// 日志级别。
	Level *string `json:"level,omitempty"`
}

func (o LogList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogList struct{}"
	}

	return strings.Join([]string{"LogList", string(data)}, " ")
}
