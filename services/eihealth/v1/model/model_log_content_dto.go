package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LogContentDto struct {

	// 作业日志采集时间
	CollectTime *string `json:"collect_time,omitempty"`

	// 作业日志内容
	Content *string `json:"content,omitempty"`
}

func (o LogContentDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogContentDto struct{}"
	}

	return strings.Join([]string{"LogContentDto", string(data)}, " ")
}
