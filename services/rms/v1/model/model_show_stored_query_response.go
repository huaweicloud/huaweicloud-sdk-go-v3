package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowStoredQueryResponse struct {

	// ResourceQL ID
	Id *string `json:"id,omitempty" xml:"id"`

	// ResourceQL 名字
	Name *string `json:"name,omitempty" xml:"name"`

	// ResourceQL 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// ResourceQL 表达式
	Expression *string `json:"expression,omitempty" xml:"expression"`

	// ResourceQL 创建时间
	Created *string `json:"created,omitempty" xml:"created"`

	// ResourceQL 更新时间
	Updated        *string `json:"updated,omitempty" xml:"updated"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowStoredQueryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStoredQueryResponse struct{}"
	}

	return strings.Join([]string{"ShowStoredQueryResponse", string(data)}, " ")
}
