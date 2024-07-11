package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckCanCreateResponse Response Object
type CheckCanCreateResponse struct {
	Result *CheckCanCreateResponseBodyResult `json:"result,omitempty"`

	// 请求成功失败状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CheckCanCreateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckCanCreateResponse struct{}"
	}

	return strings.Join([]string{"CheckCanCreateResponse", string(data)}, " ")
}
