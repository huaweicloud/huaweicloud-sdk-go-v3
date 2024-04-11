package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllStandardsResponse Response Object
type ListAllStandardsResponse struct {
	Data           *StandardSearchResultData `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListAllStandardsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllStandardsResponse struct{}"
	}

	return strings.Join([]string{"ListAllStandardsResponse", string(data)}, " ")
}
