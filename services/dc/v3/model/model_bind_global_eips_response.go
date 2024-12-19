package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindGlobalEipsResponse Response Object
type BindGlobalEipsResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	GlobalEips     *ListBindingGeip `json:"global_eips,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o BindGlobalEipsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindGlobalEipsResponse struct{}"
	}

	return strings.Join([]string{"BindGlobalEipsResponse", string(data)}, " ")
}
