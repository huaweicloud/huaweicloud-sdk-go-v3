package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTtsOnceCodeResponse Response Object
type CreateTtsOnceCodeResponse struct {

	// 获取到的一次性token
	OnceCode       *string `json:"once_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTtsOnceCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTtsOnceCodeResponse struct{}"
	}

	return strings.Join([]string{"CreateTtsOnceCodeResponse", string(data)}, " ")
}
