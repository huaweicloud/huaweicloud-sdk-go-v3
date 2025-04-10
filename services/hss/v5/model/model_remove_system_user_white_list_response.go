package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveSystemUserWhiteListResponse Response Object
type RemoveSystemUserWhiteListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveSystemUserWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveSystemUserWhiteListResponse struct{}"
	}

	return strings.Join([]string{"RemoveSystemUserWhiteListResponse", string(data)}, " ")
}
