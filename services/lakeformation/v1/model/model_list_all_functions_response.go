package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllFunctionsResponse Response Object
type ListAllFunctionsResponse struct {
	PageInfo *PagedInfo `json:"page_info,omitempty"`

	Functions      *[]Function `json:"functions,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListAllFunctionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllFunctionsResponse struct{}"
	}

	return strings.Join([]string{"ListAllFunctionsResponse", string(data)}, " ")
}
