package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnlockNodepoolNodeScaleDownResponse Response Object
type UnlockNodepoolNodeScaleDownResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UnlockNodepoolNodeScaleDownResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnlockNodepoolNodeScaleDownResponse struct{}"
	}

	return strings.Join([]string{"UnlockNodepoolNodeScaleDownResponse", string(data)}, " ")
}
