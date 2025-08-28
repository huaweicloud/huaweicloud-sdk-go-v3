package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateL7PolicyResponse Response Object
type CreateL7PolicyResponse struct {

	// **参数解释**：请求ID。  **约束限制**：不涉及  **取值范围**：由数字、小写字母和中划线（-）组成的字符串，自动生成。  **默认取值**：不涉及
	RequestId *string `json:"request_id,omitempty"`

	L7policy       *L7Policy `json:"l7policy,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateL7PolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7PolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateL7PolicyResponse", string(data)}, " ")
}
