package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LockNodepoolNodeScaleDownResponse Response Object
type LockNodepoolNodeScaleDownResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o LockNodepoolNodeScaleDownResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LockNodepoolNodeScaleDownResponse struct{}"
	}

	return strings.Join([]string{"LockNodepoolNodeScaleDownResponse", string(data)}, " ")
}
