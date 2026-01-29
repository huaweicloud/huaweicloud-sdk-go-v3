package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserPrivilegesResponse Response Object
type ListUserPrivilegesResponse struct {

	// **参数解释**：  请求成功、请求失败的状态。 **取值范围**： success: 请求成功。 error: 请求失败。
	Status *string `json:"status,omitempty"`

	// **参数解释**：  请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线(-)组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	Result         *PrivilegesResponseV5Result `json:"result,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListUserPrivilegesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserPrivilegesResponse struct{}"
	}

	return strings.Join([]string{"ListUserPrivilegesResponse", string(data)}, " ")
}
