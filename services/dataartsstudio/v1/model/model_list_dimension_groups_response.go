package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDimensionGroupsResponse Response Object
type ListDimensionGroupsResponse struct {
	Data           *ListDimensionGroupsResultData `json:"data,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListDimensionGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDimensionGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListDimensionGroupsResponse", string(data)}, " ")
}
