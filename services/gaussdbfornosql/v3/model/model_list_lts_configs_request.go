package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLtsConfigsRequest Request Object
type ListLtsConfigsRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 索引位置偏移量，表示从第一条数据偏移offset条数据后开始查询。取值必须为数字，不能为负数。默认取0值，表示从第一条数据开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。取值范围：1~100，必须为数字。不传该参数时，默认查询前100条实例信息。
	Limit *int32 `json:"limit,omitempty"`

	// 根据实例ID精确搜索。
	InstanceId *string `json:"instance_id,omitempty"`

	// 根据实例名称模糊搜索。
	InstanceName *string `json:"instance_name,omitempty"`

	// 根据企业项目ID精确搜索。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListLtsConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListLtsConfigsRequest", string(data)}, " ")
}
