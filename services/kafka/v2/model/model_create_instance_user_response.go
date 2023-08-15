package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceUserResponse Response Object
type CreateInstanceUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateInstanceUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceUserResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceUserResponse", string(data)}, " ")
}
