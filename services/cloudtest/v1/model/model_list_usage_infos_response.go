package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsageInfosResponse Response Object
type ListUsageInfosResponse struct {
	Value          *[]UsageInfos `json:"value,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListUsageInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsageInfosResponse struct{}"
	}

	return strings.Join([]string{"ListUsageInfosResponse", string(data)}, " ")
}
