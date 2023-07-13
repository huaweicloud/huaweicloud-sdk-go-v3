package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDcDsReqDto 修改数据源配置请求结构体
type UpdateDcDsReqDto struct {

	// 数据源的连接及采集信息
	Config *interface{} `json:"config,omitempty"`

	// 采集数据源名称，允许中、数字、英文大小写、下划线、中划线
	Name *string `json:"name,omitempty"`

	// 质量上报开关，不携带或值不为true，默认为false
	QualityReport *bool `json:"quality_report,omitempty"`
}

func (o UpdateDcDsReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDcDsReqDto struct{}"
	}

	return strings.Join([]string{"UpdateDcDsReqDto", string(data)}, " ")
}
