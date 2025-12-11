package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeFilePathWhiteListsResponse Response Object
type ChangeFilePathWhiteListsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeFilePathWhiteListsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeFilePathWhiteListsResponse struct{}"
	}

	return strings.Join([]string{"ChangeFilePathWhiteListsResponse", string(data)}, " ")
}
