package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyParameterGroupRequest Request Object
type CopyParameterGroupRequest struct {

	// **参数解释**：  参数组ID。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，长度为36个字符。  **默认取值**：  不涉及。
	ConfigId string `json:"config_id"`

	Body *ConfigurationCopyReqV3 `json:"body,omitempty"`
}

func (o CopyParameterGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyParameterGroupRequest struct{}"
	}

	return strings.Join([]string{"CopyParameterGroupRequest", string(data)}, " ")
}
