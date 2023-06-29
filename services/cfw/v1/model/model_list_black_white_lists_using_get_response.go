package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBlackWhiteListsUsingGetResponse Response Object
type ListBlackWhiteListsUsingGetResponse struct {
	Data           *BlackWhiteListResponseData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListBlackWhiteListsUsingGetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBlackWhiteListsUsingGetResponse struct{}"
	}

	return strings.Join([]string{"ListBlackWhiteListsUsingGetResponse", string(data)}, " ")
}
