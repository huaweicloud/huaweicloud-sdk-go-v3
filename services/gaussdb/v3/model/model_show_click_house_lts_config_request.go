package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClickHouseLtsConfigRequest Request Object
type ShowClickHouseLtsConfigRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// ClickHouse实例ID，严格匹配UUID规则。
	InstanceId *string `json:"instance_id,omitempty"`

	// 企业项目id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 查询记录数，默认10。不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`

	// 索引位置，偏移量，默认0。从第一条数据偏移offset条数据后开始查询（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowClickHouseLtsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClickHouseLtsConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowClickHouseLtsConfigRequest", string(data)}, " ")
}
