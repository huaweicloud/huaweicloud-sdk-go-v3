package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ReadConfigResp struct {

	// 图片播报配置
	ImageReadConfigs *[]ImageReadConfigResp `json:"image_read_configs,omitempty"`

	// ppt播报配置
	PptReadConfigs *[]PptReadConfigResp `json:"ppt_read_configs,omitempty"`

	// 播报选项： 0：纯文本播报 1：图片播报 2：ppt播报 默认：0 配置哪项会校验哪项是否为空
	ReadType int32 `json:"read_type"`

	// 纯文本播报内容。 换行符会按400ms的静音进行分割
	ReadContent *string `json:"read_content,omitempty"`

	// 0：左 1：中 2：右
	CharacterPosition int32 `json:"character_position"`

	// read_content 每段播报时间
	ReadContentParagraghTimes *[]int32 `json:"read_content_paragragh_times,omitempty"`
}

func (o ReadConfigResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReadConfigResp struct{}"
	}

	return strings.Join([]string{"ReadConfigResp", string(data)}, " ")
}
