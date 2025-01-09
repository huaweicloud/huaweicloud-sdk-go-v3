package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDisableAppResponse Response Object
type BatchDisableAppResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDisableAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisableAppResponse struct{}"
	}

	return strings.Join([]string{"BatchDisableAppResponse", string(data)}, " ")
}
