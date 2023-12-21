package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UninstallCbhEipResponse Response Object
type UninstallCbhEipResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UninstallCbhEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UninstallCbhEipResponse struct{}"
	}

	return strings.Join([]string{"UninstallCbhEipResponse", string(data)}, " ")
}
