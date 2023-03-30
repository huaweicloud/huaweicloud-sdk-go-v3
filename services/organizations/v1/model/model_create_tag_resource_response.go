package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateTagResourceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateTagResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagResourceResponse struct{}"
	}

	return strings.Join([]string{"CreateTagResourceResponse", string(data)}, " ")
}
