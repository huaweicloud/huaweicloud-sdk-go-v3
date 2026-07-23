package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BusinessCodeVo struct {

	// **参数解释：**  业务编码，根据模型配置的业务编码生成规则自动生成的流水。  **取值范围：**  不涉及。
	Code *string `json:"code,omitempty"`
}

func (o BusinessCodeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessCodeVo struct{}"
	}

	return strings.Join([]string{"BusinessCodeVo", string(data)}, " ")
}
