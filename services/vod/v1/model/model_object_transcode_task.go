package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObjectTranscodeTask struct {

	// 转码模板ID
	TransTemplateId string `json:"trans_template_id"`

	// 转码输出文件名
	OutputName *string `json:"output_name,omitempty"`

	Output *ObsInfo `json:"output,omitempty"`
}

func (o ObjectTranscodeTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectTranscodeTask struct{}"
	}

	return strings.Join([]string{"ObjectTranscodeTask", string(data)}, " ")
}
