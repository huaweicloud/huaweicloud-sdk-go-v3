package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataSummaryRsp 数据对象
type DataSummaryRsp struct {

	// 对象全路径（项目名称:/路径）
	Path *string `json:"path,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	Type *PathType `json:"type,omitempty"`

	// 大小
	Size *int64 `json:"size,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 可操作标记
	AllowedOperate *bool `json:"allowed_operate,omitempty"`

	// 可删除标记
	Deletable *bool `json:"deletable,omitempty"`
}

func (o DataSummaryRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataSummaryRsp struct{}"
	}

	return strings.Join([]string{"DataSummaryRsp", string(data)}, " ")
}
