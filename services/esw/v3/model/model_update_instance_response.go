package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceResponse Response Object
type UpdateInstanceResponse struct {

	// - 参数解释：请求的唯一标识。 - 约束限制：UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	RequestId *string `json:"request_id,omitempty"`

	Instance       *Instance `json:"instance,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceResponse", string(data)}, " ")
}
