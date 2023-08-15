package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddOtTemplatesResponse Response Object
type AddOtTemplatesResponse struct {

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

func (o AddOtTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddOtTemplatesResponse struct{}"
	}

	return strings.Join([]string{"AddOtTemplatesResponse", string(data)}, " ")
}
