package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateResourceResponse struct {
	Body           *string `json:"body,omitempty" xml:"body"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceResponse struct{}"
	}

	return strings.Join([]string{"CreateResourceResponse", string(data)}, " ")
}
