package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDcDsReqDto 创建数据源配置请求结构体
type CreateDcDsReqDto struct {

	// 采集数据源id，节点下唯一
	DsId string `json:"ds_id"`

	// 数据源的连接及采集信息
	Config *interface{} `json:"config"`

	// 采集数据源名称，允许中、数字、英文大小写、下划线、中划线
	Name string `json:"name"`

	// 模块id
	ModuleId string `json:"module_id"`

	// 模板id，节点下唯一
	TplId string `json:"tpl_id"`

	// 质量上报开关，不携带或值不为true，默认为false
	QualityReport *bool `json:"quality_report,omitempty"`
}

func (o CreateDcDsReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDcDsReqDto struct{}"
	}

	return strings.Join([]string{"CreateDcDsReqDto", string(data)}, " ")
}
