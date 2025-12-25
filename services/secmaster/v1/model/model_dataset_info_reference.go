package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatasetInfoReference 数据源名称
type DatasetInfoReference struct {

	// 云服务描述
	CsvcDisplay string `json:"csvc_display"`

	// 数据源描述
	SourceDisplay string `json:"source_display"`

	// 链接
	Link *string `json:"link,omitempty"`

	// 云服务帮助说明
	CsvcHelp *string `json:"csvc_help,omitempty"`

	// 数据源帮助说明
	SourceHelp *string `json:"source_help,omitempty"`
}

func (o DatasetInfoReference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatasetInfoReference struct{}"
	}

	return strings.Join([]string{"DatasetInfoReference", string(data)}, " ")
}
