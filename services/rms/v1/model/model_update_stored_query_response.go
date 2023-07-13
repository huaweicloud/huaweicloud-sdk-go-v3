package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStoredQueryResponse Response Object
type UpdateStoredQueryResponse struct {

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

	// ResourceQL 更新时间
	Updated        *string `json:"updated,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateStoredQueryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStoredQueryResponse struct{}"
	}

	return strings.Join([]string{"UpdateStoredQueryResponse", string(data)}, " ")
}
