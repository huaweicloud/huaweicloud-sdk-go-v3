package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListModifyHistoryRequest Request Object
type ListModifyHistoryRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 参数组ID。  通过调用[查询实例详情信息](https://support.huaweicloud.com/api-taurusdb/ShowGaussMySqlInstanceInfo.html)接口获取。  请求响应成功后在响应消息体中包含的“configuration_id”的值即为configuration_id值。
	ConfigurationId string `json:"configuration_id"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListModifyHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListModifyHistoryRequest struct{}"
	}

	return strings.Join([]string{"ListModifyHistoryRequest", string(data)}, " ")
}
