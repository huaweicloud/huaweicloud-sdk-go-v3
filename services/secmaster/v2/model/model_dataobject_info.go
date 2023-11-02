package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataobjectInfo 数据对象详情
type DataobjectInfo struct {

	// ID值
	Id *string `json:"id,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 数据类ID
	DataclassId *string `json:"dataclass_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 数据内容
	Content *string `json:"content,omitempty"`
}

func (o DataobjectInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataobjectInfo struct{}"
	}

	return strings.Join([]string{"DataobjectInfo", string(data)}, " ")
}
