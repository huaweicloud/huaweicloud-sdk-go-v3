package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupplementdataResponse Response Object
type ListSupplementdataResponse struct {
	Msg *string `json:"msg,omitempty"`

	Rows *[]SupplementDataResp `json:"rows,omitempty"`

	Success *bool `json:"success,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSupplementdataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupplementdataResponse struct{}"
	}

	return strings.Join([]string{"ListSupplementdataResponse", string(data)}, " ")
}
