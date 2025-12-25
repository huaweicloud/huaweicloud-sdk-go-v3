package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGenericActionResponse Response Object
type CreateGenericActionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateGenericActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGenericActionResponse struct{}"
	}

	return strings.Join([]string{"CreateGenericActionResponse", string(data)}, " ")
}
