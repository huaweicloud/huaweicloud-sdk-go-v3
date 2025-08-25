package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCcspInstanceInfoRequest Request Object
type ShowCcspInstanceInfoRequest struct {

	// 指定查询返回记录条数，默认值10
	PageSize *int32 `json:"page_size,omitempty"`

	// 实例名称
	Name *string `json:"name,omitempty"`

	// 索引位置，从page_num指定的下一条数据开始查询默认值为0
	PageNum *int32 `json:"page_num,omitempty"`

	// 排序属性，目前支持以下属性： - **create_time** : 实例创建时间（默认）
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方向，支持以下值： - **DESC** : 降序（默认） - **ASC** : 升序
	SortDir *string `json:"sort_dir,omitempty"`

	// 密码服务集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 实例的服务状态: - **enable** : 启用； - **disable** : 停用
	Status *int32 `json:"status,omitempty"`

	// 实例健康状态： - **true** : 正常； - **false** : 异常
	IsNormal *bool `json:"is_normal,omitempty"`
}

func (o ShowCcspInstanceInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCcspInstanceInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowCcspInstanceInfoRequest", string(data)}, " ")
}
