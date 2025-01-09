package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchEnableAppResponse Response Object
type BatchEnableAppResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchEnableAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnableAppResponse struct{}"
	}

	return strings.Join([]string{"BatchEnableAppResponse", string(data)}, " ")
}
