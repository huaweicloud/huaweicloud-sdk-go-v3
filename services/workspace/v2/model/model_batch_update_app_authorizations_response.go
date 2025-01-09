package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateAppAuthorizationsResponse Response Object
type BatchUpdateAppAuthorizationsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUpdateAppAuthorizationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAppAuthorizationsResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateAppAuthorizationsResponse", string(data)}, " ")
}
