package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMyInfoResponse Response Object
type UpdateMyInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateMyInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMyInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateMyInfoResponse", string(data)}, " ")
}
