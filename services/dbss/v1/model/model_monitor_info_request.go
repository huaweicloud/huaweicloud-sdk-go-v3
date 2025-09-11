package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MonitorInfoRequest struct {
	Time *Duration `json:"time"`
}

func (o MonitorInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MonitorInfoRequest struct{}"
	}

	return strings.Join([]string{"MonitorInfoRequest", string(data)}, " ")
}
