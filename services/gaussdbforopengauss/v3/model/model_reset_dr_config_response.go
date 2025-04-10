package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetDrConfigResponse Response Object
type ResetDrConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetDrConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetDrConfigResponse struct{}"
	}

	return strings.Join([]string{"ResetDrConfigResponse", string(data)}, " ")
}
