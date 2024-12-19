package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnbindGlobalEipsResponse Response Object
type UnbindGlobalEipsResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	GlobalEips     *ListBindingGeip `json:"global_eips,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UnbindGlobalEipsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindGlobalEipsResponse struct{}"
	}

	return strings.Join([]string{"UnbindGlobalEipsResponse", string(data)}, " ")
}
