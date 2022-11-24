package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateDataStoreResponse struct {

	// 存储 ID
	DataStoreId *string `json:"data_store_id,omitempty"`

	// 存储名称
	Name *string `json:"name,omitempty"`

	// 存储 ID
	GroupId *string `json:"group_id,omitempty"`

	// 标签
	Tags *[]Tag `json:"tags,omitempty"`

	// 指标
	Metrics *[]Metric `json:"metrics,omitempty"`

	// 属性
	Properties *[]Property `json:"properties,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新时间
	ModifiedTime   *string `json:"modified_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDataStoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataStoreResponse struct{}"
	}

	return strings.Join([]string{"UpdateDataStoreResponse", string(data)}, " ")
}
