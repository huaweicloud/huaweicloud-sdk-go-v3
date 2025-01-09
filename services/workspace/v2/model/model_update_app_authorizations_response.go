package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppAuthorizationsResponse Response Object
type UpdateAppAuthorizationsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAppAuthorizationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppAuthorizationsResponse struct{}"
	}

	return strings.Join([]string{"UpdateAppAuthorizationsResponse", string(data)}, " ")
}
