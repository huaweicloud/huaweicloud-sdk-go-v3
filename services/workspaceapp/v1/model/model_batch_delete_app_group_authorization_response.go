package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAppGroupAuthorizationResponse Response Object
type BatchDeleteAppGroupAuthorizationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteAppGroupAuthorizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAppGroupAuthorizationResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAppGroupAuthorizationResponse", string(data)}, " ")
}
