package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompareVersionRespVo struct {

	// **参数解释：**  基础版本对象。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	BasicVersion *interface{} `json:"basicVersion,omitempty"`

	// **参数解释：**  当前版本对象。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CorrelationVersion *interface{} `json:"correlationVersion,omitempty"`
}

func (o CompareVersionRespVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareVersionRespVo struct{}"
	}

	return strings.Join([]string{"CompareVersionRespVo", string(data)}, " ")
}
