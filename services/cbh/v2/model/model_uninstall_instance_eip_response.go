package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UninstallInstanceEipResponse Response Object
type UninstallInstanceEipResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UninstallInstanceEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UninstallInstanceEipResponse struct{}"
	}

	return strings.Join([]string{"UninstallInstanceEipResponse", string(data)}, " ")
}
