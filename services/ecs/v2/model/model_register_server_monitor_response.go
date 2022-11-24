package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RegisterServerMonitorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RegisterServerMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterServerMonitorResponse struct{}"
	}

	return strings.Join([]string{"RegisterServerMonitorResponse", string(data)}, " ")
}
