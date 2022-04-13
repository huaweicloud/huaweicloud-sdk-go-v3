package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateVisionActiveCodeResponse struct {
	// 激活码

	ActiveCode     *string `json:"activeCode,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateVisionActiveCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVisionActiveCodeResponse struct{}"
	}

	return strings.Join([]string{"CreateVisionActiveCodeResponse", string(data)}, " ")
}
