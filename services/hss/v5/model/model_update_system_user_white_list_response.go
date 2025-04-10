package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSystemUserWhiteListResponse Response Object
type UpdateSystemUserWhiteListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSystemUserWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSystemUserWhiteListResponse struct{}"
	}

	return strings.Join([]string{"UpdateSystemUserWhiteListResponse", string(data)}, " ")
}
