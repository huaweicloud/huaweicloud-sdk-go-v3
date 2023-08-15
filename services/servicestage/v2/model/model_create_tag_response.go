package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTagResponse Response Object
type CreateTagResponse struct {

	// tag标签名称。
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagResponse struct{}"
	}

	return strings.Join([]string{"CreateTagResponse", string(data)}, " ")
}
