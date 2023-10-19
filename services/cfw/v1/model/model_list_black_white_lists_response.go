package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBlackWhiteListsResponse Response Object
type ListBlackWhiteListsResponse struct {
	Data           *BlackWhiteListResponseData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListBlackWhiteListsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBlackWhiteListsResponse struct{}"
	}

	return strings.Join([]string{"ListBlackWhiteListsResponse", string(data)}, " ")
}
