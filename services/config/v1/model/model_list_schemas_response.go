package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSchemasResponse Response Object
type ListSchemasResponse struct {

	// schemas 接口.
	Value *[]ResourceSchemaResponse `json:"value,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSchemasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSchemasResponse struct{}"
	}

	return strings.Join([]string{"ListSchemasResponse", string(data)}, " ")
}
