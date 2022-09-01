package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetDataStore struct {

	// 存储 ID
	DataStoreId *string `json:"data_store_id,omitempty" xml:"data_store_id"`

	// 存储名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 存储 ID
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 标签
	Tags *[]Tag `json:"tags,omitempty" xml:"tags"`

	// 指标
	Metrics *[]Metric `json:"metrics,omitempty" xml:"metrics"`

	// 属性
	Properties *[]Property `json:"properties,omitempty" xml:"properties"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty" xml:"created_time"`

	// 更新时间
	ModifiedTime *string `json:"modified_time,omitempty" xml:"modified_time"`
}

func (o GetDataStore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDataStore struct{}"
	}

	return strings.Join([]string{"GetDataStore", string(data)}, " ")
}
