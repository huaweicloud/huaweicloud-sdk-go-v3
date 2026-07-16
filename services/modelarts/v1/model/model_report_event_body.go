package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReportEventBody struct {

	// 事件上报的类型
	Type *string `json:"type,omitempty"`
}

func (o ReportEventBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportEventBody struct{}"
	}

	return strings.Join([]string{"ReportEventBody", string(data)}, " ")
}
