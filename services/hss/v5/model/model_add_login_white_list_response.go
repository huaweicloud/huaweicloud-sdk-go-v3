package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddLoginWhiteListResponse Response Object
type AddLoginWhiteListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddLoginWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddLoginWhiteListResponse struct{}"
	}

	return strings.Join([]string{"AddLoginWhiteListResponse", string(data)}, " ")
}
