package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskOutPut struct {
	ObsInfo *ObsInfo `json:"obs_info,omitempty"`

	MetaData *ObjectMetaData `json:"meta_data,omitempty"`

	// 输出文件播放地址
	FileUrl *string `json:"file_url,omitempty"`
}

func (o TaskOutPut) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskOutPut struct{}"
	}

	return strings.Join([]string{"TaskOutPut", string(data)}, " ")
}
