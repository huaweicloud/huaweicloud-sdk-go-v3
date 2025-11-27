package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeSensitiveInfoRequestInfo **参数解释**: 操作的事件
type ChangeSensitiveInfoRequestInfo struct {

	// **参数解释**: 敏感信息编号 **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	SensitiveInfoId string `json:"sensitive_info_id"`

	// **参数解释**: 操作敏感信息详情，处理方式 **约束限制**: 不涉及 **取值范围**: - ignore：忽略 - do_not_ignore：取消忽略  **默认取值**: do_not_ignore
	OperateType *string `json:"operate_type,omitempty"`
}

func (o ChangeSensitiveInfoRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSensitiveInfoRequestInfo struct{}"
	}

	return strings.Join([]string{"ChangeSensitiveInfoRequestInfo", string(data)}, " ")
}
