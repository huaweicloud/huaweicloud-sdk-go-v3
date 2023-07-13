package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOtTemplatesReqDto 创建OT数采模板请求体
type CreateOtTemplatesReqDto struct {

	// 模板id
	TplId string `json:"tpl_id"`

	// 模板名称，允许中、数字、英文大小写、下划线、中划线
	Name string `json:"name"`

	// 描述
	Description string `json:"description"`

	// 数据源元数据
	DatasourceMeta *interface{} `json:"datasource_meta"`

	// 点位表元数据
	PointMeta *interface{} `json:"point_meta"`
}

func (o CreateOtTemplatesReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOtTemplatesReqDto struct{}"
	}

	return strings.Join([]string{"CreateOtTemplatesReqDto", string(data)}, " ")
}
