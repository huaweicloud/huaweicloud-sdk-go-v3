package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveLoginWhiteListResponse Response Object
type RemoveLoginWhiteListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveLoginWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveLoginWhiteListResponse struct{}"
	}

	return strings.Join([]string{"RemoveLoginWhiteListResponse", string(data)}, " ")
}
