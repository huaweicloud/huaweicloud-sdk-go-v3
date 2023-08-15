package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionsResponse Response Object
type ListFunctionsResponse struct {
	PageInfo *PagedInfo `json:"page_info,omitempty"`

	Functions      *[]Function `json:"functions,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListFunctionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionsResponse struct{}"
	}

	return strings.Join([]string{"ListFunctionsResponse", string(data)}, " ")
}
