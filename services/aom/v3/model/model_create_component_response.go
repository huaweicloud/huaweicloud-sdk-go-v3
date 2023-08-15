package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateComponentResponse Response Object
type CreateComponentResponse struct {

	// 对象id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateComponentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentResponse struct{}"
	}

	return strings.Join([]string{"CreateComponentResponse", string(data)}, " ")
}
