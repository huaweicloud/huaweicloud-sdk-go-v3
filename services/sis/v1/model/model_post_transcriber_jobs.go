package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type PostTranscriberJobs struct {
	Config *TranscriberConfig `json:"config"`
	// 存放在OBS的音频文件路径。OBS的region要和请求服务的region保持一致，region不一致则OBS不可用，即使obs是公开访问权限。

	DataUrl string `json:"data_url"`
}

func (o PostTranscriberJobs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostTranscriberJobs struct{}"
	}

	return strings.Join([]string{"PostTranscriberJobs", string(data)}, " ")
}
