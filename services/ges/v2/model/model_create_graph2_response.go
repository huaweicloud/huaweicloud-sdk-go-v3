package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGraph2Response Response Object
type CreateGraph2Response struct {

	// 图ID。
	Id *string `json:"id,omitempty"`

	// 图名称。
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGraph2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGraph2Response struct{}"
	}

	return strings.Join([]string{"CreateGraph2Response", string(data)}, " ")
}
