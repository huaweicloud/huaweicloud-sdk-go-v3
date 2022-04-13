package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateMpeCallBackResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateMpeCallBackResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMpeCallBackResponse struct{}"
	}

	return strings.Join([]string{"CreateMpeCallBackResponse", string(data)}, " ")
}
