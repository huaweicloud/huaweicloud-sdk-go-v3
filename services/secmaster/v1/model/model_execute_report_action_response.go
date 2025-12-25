package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteReportActionResponse Response Object
type ExecuteReportActionResponse struct {

	// **参数解释**: 错误码 **取值范围**: 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**: 错误描述 **取值范围**: 不涉及
	Message *string `json:"message,omitempty"`

	// 导出文件
	File           *interface{} `json:"file,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ExecuteReportActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteReportActionResponse struct{}"
	}

	return strings.Join([]string{"ExecuteReportActionResponse", string(data)}, " ")
}
