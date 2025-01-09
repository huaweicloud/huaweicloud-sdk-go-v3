package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchChangeServerMaintainModeResponse Response Object
type BatchChangeServerMaintainModeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchChangeServerMaintainModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeServerMaintainModeResponse struct{}"
	}

	return strings.Join([]string{"BatchChangeServerMaintainModeResponse", string(data)}, " ")
}
