package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssetModelModRequest struct {

	// 模型显示名称，正则：\"^[\\\\u4E00-\\\\u9FA5A-Za-z0-9_@#.-]{1,64}$\"
	DisplayName *string `json:"display_name,omitempty"`

	// 属性集，最多200个
	Properties *[]PropertyModelRequest `json:"properties,omitempty"`

	// 分析任务集，最多50个
	Analyses *[]AnalysisModelRequest `json:"analyses,omitempty"`
}

func (o AssetModelModRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetModelModRequest struct{}"
	}

	return strings.Join([]string{"AssetModelModRequest", string(data)}, " ")
}
