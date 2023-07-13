package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomPropsTaskResponse Response Object
type CreateCustomPropsTaskResponse struct {

	// 自定义属性ID
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCustomPropsTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomPropsTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateCustomPropsTaskResponse", string(data)}, " ")
}
