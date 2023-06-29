package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMetadata2Response Response Object
type CreateMetadata2Response struct {

	// 元数据ID。
	Id *string `json:"id,omitempty"`

	// 元数据名字。
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMetadata2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetadata2Response struct{}"
	}

	return strings.Join([]string{"CreateMetadata2Response", string(data)}, " ")
}
