package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateComparePolicyResponse Response Object
type UpdateComparePolicyResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateComparePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComparePolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateComparePolicyResponse", string(data)}, " ")
}
