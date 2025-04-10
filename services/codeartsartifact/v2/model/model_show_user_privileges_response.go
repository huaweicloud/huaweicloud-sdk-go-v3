package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserPrivilegesResponse Response Object
type ShowUserPrivilegesResponse struct {

	// **参数解释**:  请求成功、失败状态。 **取值范围**: success: 请求成功。 error: 请求失败。
	Status *string `json:"status,omitempty"`

	// **参数解释**:  请求id，当前请求唯一标识。 **取值范围**: 数字及中划线(-)组成的字符串。
	TraceId *string `json:"traceId,omitempty"`

	Result         *PrivilegesResponseResult `json:"result,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ShowUserPrivilegesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserPrivilegesResponse struct{}"
	}

	return strings.Join([]string{"ShowUserPrivilegesResponse", string(data)}, " ")
}
