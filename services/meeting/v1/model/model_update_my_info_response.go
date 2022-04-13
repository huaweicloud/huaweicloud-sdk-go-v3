package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
