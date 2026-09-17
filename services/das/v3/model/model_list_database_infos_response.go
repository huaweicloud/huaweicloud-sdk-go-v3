package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseInfosResponse Response Object
type ListDatabaseInfosResponse struct {
	Body           *[]DatabaseUsageInfoResp `json:"body,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListDatabaseInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseInfosResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseInfosResponse", string(data)}, " ")
}
