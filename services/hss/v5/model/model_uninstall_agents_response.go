package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UninstallAgentsResponse Response Object
type UninstallAgentsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UninstallAgentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UninstallAgentsResponse struct{}"
	}

	return strings.Join([]string{"UninstallAgentsResponse", string(data)}, " ")
}
