package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppResourceLimitResponse Response Object
type ShowAppResourceLimitResponse struct {

	// **参数解释**： 允许创建的应用数量。 **取值范围**： - 0：不支持创建。 - 大于0：允许创建的应用数量。
	AllowedResourceNumber *int32 `json:"allowed_resource_number,omitempty"`

	XTraceId       *string `json:"X-Trace-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAppResourceLimitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppResourceLimitResponse struct{}"
	}

	return strings.Join([]string{"ShowAppResourceLimitResponse", string(data)}, " ")
}
