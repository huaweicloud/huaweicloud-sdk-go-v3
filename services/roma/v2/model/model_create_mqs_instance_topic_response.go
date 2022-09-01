package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateMqsInstanceTopicResponse struct {

	// topic名称。
	Name           *string `json:"name,omitempty" xml:"name"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMqsInstanceTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMqsInstanceTopicResponse struct{}"
	}

	return strings.Join([]string{"CreateMqsInstanceTopicResponse", string(data)}, " ")
}
