package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssetModRequest struct {
	// 资产名称，修改资产时，null或不携带代表不修改，正则：\"^[a-zA-Z][a-zA-Z0-9_-]{0,63}$\"

	Name *string `json:"name,omitempty"`
	// 资产显示名称，修改资产时，\"\"代表配置为空、null或不携带代表不修改，正则：\"^[\\\\u4E00-\\\\u9FA5A-Za-z0-9_@#.-]{0,64}$\"

	DisplayName *string `json:"display_name,omitempty"`
	// 父资产ID，根资产的父资产ID为null，修改资产时，null或不携带代表不修改

	Parent *string `json:"parent,omitempty"`
	// 属性集，最多200个

	Properties *[]PropertyRequest `json:"properties,omitempty"`
	// 分析任务集，最多50个

	Analyses *[]AnalysisRequest `json:"analyses,omitempty"`
}

func (o AssetModRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetModRequest struct{}"
	}

	return strings.Join([]string{"AssetModRequest", string(data)}, " ")
}
