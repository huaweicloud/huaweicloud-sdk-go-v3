package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDataStore struct {

	// 存储名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 标签，更新存储时标签只可新增，不可修改或删除原有标签
	Tags *[]Tag `json:"tags,omitempty" xml:"tags"`

	// 指标
	Metrics *[]Metric `json:"metrics,omitempty" xml:"metrics"`

	// 属性，更新存储时属性只可新增，不可修改或删除原有属性
	Properties *[]Property `json:"properties,omitempty" xml:"properties"`
}

func (o UpdateDataStore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataStore struct{}"
	}

	return strings.Join([]string{"UpdateDataStore", string(data)}, " ")
}
