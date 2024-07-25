package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteByConditionVo struct {
	Condition *QueryRequestVo `json:"condition"`

	// **参数解释：**  修改人。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o DeleteByConditionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteByConditionVo struct{}"
	}

	return strings.Join([]string{"DeleteByConditionVo", string(data)}, " ")
}
