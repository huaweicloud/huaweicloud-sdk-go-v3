package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSchemasResponse struct {

	// schemas 接口.
	Value *[]ResourceSchemaResponse `json:"value,omitempty" xml:"value"`

	PageInfo       *PageInfo `json:"page_info,omitempty" xml:"page_info"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSchemasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSchemasResponse struct{}"
	}

	return strings.Join([]string{"ListSchemasResponse", string(data)}, " ")
}