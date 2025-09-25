package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeImageWhiteListResponse Response Object
type ChangeImageWhiteListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeImageWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeImageWhiteListResponse struct{}"
	}

	return strings.Join([]string{"ChangeImageWhiteListResponse", string(data)}, " ")
}
