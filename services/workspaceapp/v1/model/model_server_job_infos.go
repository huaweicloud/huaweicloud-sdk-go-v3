package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerJobInfos struct {

	// 服务器任务信息。
	Items *[]ServerJobInfo `json:"items,omitempty"`
}

func (o ServerJobInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerJobInfos struct{}"
	}

	return strings.Join([]string{"ServerJobInfos", string(data)}, " ")
}
