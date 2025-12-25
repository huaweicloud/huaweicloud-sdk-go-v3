package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LtsResponseVoLtsInfos struct {

	// 日志组id
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 组名称
	LogGroupName *string `json:"log_group_name,omitempty"`

	// 流列表
	Streams *[]LtsResponseVoStreams `json:"streams,omitempty"`
}

func (o LtsResponseVoLtsInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsResponseVoLtsInfos struct{}"
	}

	return strings.Join([]string{"LtsResponseVoLtsInfos", string(data)}, " ")
}
