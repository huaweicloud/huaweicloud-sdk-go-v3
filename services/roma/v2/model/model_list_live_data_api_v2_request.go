package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListLiveDataApiV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 后端API归属的集成应用编号
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 后端API归属的集成应用名称
	AppName *string `json:"app_name,omitempty" xml:"app_name"`

	// 后端API名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 后端API状态，支持1，3，4，分别表示待开发，开发中和已部署状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 后端API请求路径
	Path *string `json:"path,omitempty" xml:"path"`

	// 指定需要精确匹配查找的参数名称，多个参数需要支持精确匹配时参数之间使用“,”隔开。  当前支持name，path，status。
	PreciseSearch *string `json:"precise_search,omitempty" xml:"precise_search"`
}

func (o ListLiveDataApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLiveDataApiV2Request struct{}"
	}

	return strings.Join([]string{"ListLiveDataApiV2Request", string(data)}, " ")
}
