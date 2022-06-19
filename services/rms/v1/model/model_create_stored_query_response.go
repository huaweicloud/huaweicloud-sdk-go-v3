package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateStoredQueryResponse struct {

	// ResourceQL ID
	Id *string `json:"id,omitempty"`

	// ResourceQL 名字
	Name *string `json:"name,omitempty"`

	// ResourceQL 描述
	Description *string `json:"description,omitempty"`

	// ResourceQL 表达式
	Expression *string `json:"expression,omitempty"`

	// ResourceQL 创建时间
	Created *string `json:"created,omitempty"`

	// ResouerceQL 更新时间
	Updated        *string `json:"updated,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateStoredQueryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStoredQueryResponse struct{}"
	}

	return strings.Join([]string{"CreateStoredQueryResponse", string(data)}, " ")
}
