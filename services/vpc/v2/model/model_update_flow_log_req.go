package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type UpdateFlowLogReq struct {

	// 功能说明：流日志名称 取值范围：0-64个字符，支持数字、字母、中文、_（下划线）、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// 功能说明：流日志描述 取值范围：0-255个字符，不能包含“<”和“>”
	Description *string `json:"description,omitempty"`

	// 功能说明：流日志管理 取值范围：若为true，表明开启流日志。若为false，则关闭流日志。
	AdminState *bool `json:"admin_state,omitempty"`
}

func (o UpdateFlowLogReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlowLogReq struct{}"
	}

	return strings.Join([]string{"UpdateFlowLogReq", string(data)}, " ")
}
