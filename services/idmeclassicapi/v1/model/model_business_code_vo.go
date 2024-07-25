package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BusinessCodeVo struct {

	// **参数解释：**  业务编码。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Code *string `json:"code,omitempty"`
}

func (o BusinessCodeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessCodeVo struct{}"
	}

	return strings.Join([]string{"BusinessCodeVo", string(data)}, " ")
}
