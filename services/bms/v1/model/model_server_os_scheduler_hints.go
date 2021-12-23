package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerOsSchedulerHints struct {
	// 反亲和性组信息。  UUID格式。

	Group *[]string `json:"group,omitempty"`
}

func (o ServerOsSchedulerHints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerOsSchedulerHints struct{}"
	}

	return strings.Join([]string{"ServerOsSchedulerHints", string(data)}, " ")
}
