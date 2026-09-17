package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateDbDataResponse Response Object
type ValidateDbDataResponse struct {

	// **参数解释**： 校验总结果数。 **默认取值**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 校验成功结果数。 **默认取值**： 不涉及。
	Success *int32 `json:"success,omitempty"`

	// **参数解释**： 校验失败结果数。 **默认取值**： 不涉及。
	Failure *int32 `json:"failure,omitempty"`

	// **参数解释**： 校验数据类型。 **默认取值**： schema、table
	Type *string `json:"type,omitempty"`

	// **参数解释**： 校验成功结果数据。 **默认取值**： 不涉及。
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ValidateDbDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateDbDataResponse struct{}"
	}

	return strings.Join([]string{"ValidateDbDataResponse", string(data)}, " ")
}
