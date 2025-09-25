package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteImageWhiteListsResponse Response Object
type DeleteImageWhiteListsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteImageWhiteListsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteImageWhiteListsResponse struct{}"
	}

	return strings.Join([]string{"DeleteImageWhiteListsResponse", string(data)}, " ")
}
