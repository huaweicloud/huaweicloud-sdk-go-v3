package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddImageWhiteListsResponse Response Object
type AddImageWhiteListsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddImageWhiteListsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddImageWhiteListsResponse struct{}"
	}

	return strings.Join([]string{"AddImageWhiteListsResponse", string(data)}, " ")
}
