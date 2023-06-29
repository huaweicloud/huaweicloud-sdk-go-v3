package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLiveDataDataSourcesV2Request Request Object
type ListLiveDataDataSourcesV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 应用编号
	AppId string `json:"app_id"`

	// 数据源名称
	Name *string `json:"name,omitempty"`
}

func (o ListLiveDataDataSourcesV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLiveDataDataSourcesV2Request struct{}"
	}

	return strings.Join([]string{"ListLiveDataDataSourcesV2Request", string(data)}, " ")
}
