package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppDetailsRequest Request Object
type ListAppDetailsRequest struct {

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 地域 1. cn：国内 2. intl：国际
	Region *string `json:"region,omitempty"`

	// 排序方式 1. desc：降序 2. asc：升序
	SortDir *string `json:"sort_dir,omitempty"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`

	// 应用状态   1. CREATED：待上线。应用暂未创建成功，请稍候。   2. SUSPENDED：暂停。无法发起业务请求。当客户所发短信内容触发业务违规，或客户申请退订短信业务时，运营经理会将客户短信应用暂停。   3. LAUNCHED：正常。应用添加成功，可以正常使用。   4. PROCESSING：资源待分配。请联系客户经理或通过工单系统申请配置资源。
	Status *string `json:"status,omitempty"`
}

func (o ListAppDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListAppDetailsRequest", string(data)}, " ")
}
