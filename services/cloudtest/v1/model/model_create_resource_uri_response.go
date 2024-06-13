package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceUriResponse Response Object
type CreateResourceUriResponse struct {
	Value          *string `json:"value,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateResourceUriResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceUriResponse struct{}"
	}

	return strings.Join([]string{"CreateResourceUriResponse", string(data)}, " ")
}
