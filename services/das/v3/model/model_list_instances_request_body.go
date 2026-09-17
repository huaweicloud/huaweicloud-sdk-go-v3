package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesRequestBody 获取实例列表请求体
type ListInstancesRequestBody struct {

	// 数据库引擎类型
	EngineType *string `json:"engine_type,omitempty"`

	// 实例状态
	InstanceStatus *string `json:"instance_status,omitempty"`

	// 页数
	CurPage *int32 `json:"cur_page,omitempty"`

	// 页大小
	PageSize *int32 `json:"page_size,omitempty"`

	// 实例类型
	InstanceType *string `json:"instance_type,omitempty"`

	// 实例版本
	EngineVersion *string `json:"engine_version,omitempty"`

	// 历史事务是否开启
	TransactionFlag *bool `json:"transaction_flag,omitempty"`
}

func (o ListInstancesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequestBody struct{}"
	}

	return strings.Join([]string{"ListInstancesRequestBody", string(data)}, " ")
}
