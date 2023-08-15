package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCodeResponse Response Object
type CreateCodeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCodeResponse struct{}"
	}

	return strings.Join([]string{"CreateCodeResponse", string(data)}, " ")
}
