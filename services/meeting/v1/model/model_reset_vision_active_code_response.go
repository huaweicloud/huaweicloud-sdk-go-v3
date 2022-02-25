package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ResetVisionActiveCodeResponse struct {
	// 激活码

	ActiveCode     *string `json:"activeCode,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetVisionActiveCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetVisionActiveCodeResponse struct{}"
	}

	return strings.Join([]string{"ResetVisionActiveCodeResponse", string(data)}, " ")
}
