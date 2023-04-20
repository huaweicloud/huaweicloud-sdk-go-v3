package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowOtTemplateResponse struct {

	// 模板id
	TplId *string `json:"tpl_id,omitempty"`

	// 模板名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 数据源元数据
	DatasourceMeta *interface{} `json:"datasource_meta,omitempty"`

	// 点位表元数据
	PointMeta *interface{} `json:"point_meta,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最后一次修改时间
	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowOtTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOtTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowOtTemplateResponse", string(data)}, " ")
}
