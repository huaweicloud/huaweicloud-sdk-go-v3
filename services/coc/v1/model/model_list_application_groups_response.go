package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationGroupsResponse Response Object
type ListApplicationGroupsResponse struct {

	// 分组ID。
	Data           *[]GroupQueryResponseData `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListApplicationGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationGroupsResponse", string(data)}, " ")
}
