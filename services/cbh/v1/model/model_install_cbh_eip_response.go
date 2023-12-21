package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstallCbhEipResponse Response Object
type InstallCbhEipResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o InstallCbhEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallCbhEipResponse struct{}"
	}

	return strings.Join([]string{"InstallCbhEipResponse", string(data)}, " ")
}
