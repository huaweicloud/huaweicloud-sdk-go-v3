package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationCopyReqV3 请求体
type ConfigurationCopyReqV3 struct {
	CopyPara *ParaGroupCopy `json:"copy_para,omitempty"`

	// **参数解释**：  目标参数组ID。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，长度为36个字符。  **默认取值**：  不涉及。
	SourceId *string `json:"source_id,omitempty"`
}

func (o ConfigurationCopyReqV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationCopyReqV3 struct{}"
	}

	return strings.Join([]string{"ConfigurationCopyReqV3", string(data)}, " ")
}
