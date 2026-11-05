package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuthGrantResponse Response Object
type CreateAuthGrantResponse struct {

	// **参数解释**：服务名称。 **取值范围**：不涉及。
	ServiceName *string `json:"service_name,omitempty"`

	// **参数解释**：授权状态，true表示已授权，false表示未授权。 **取值范围**： - true：已授权。 - false：未授权。
	Status         *bool `json:"status,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateAuthGrantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthGrantResponse struct{}"
	}

	return strings.Join([]string{"CreateAuthGrantResponse", string(data)}, " ")
}
