package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLtsConfigsRequest Request Object
type ShowLtsConfigsRequest struct {

	// 内容类型。  取值：application/json。
	ContentType string `json:"Content-Type"`

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 查询记录数，默认值为10，最小为1，最大为100。
	Limit *string `json:"limit,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询）。
	Offset *string `json:"offset,omitempty"`
}

func (o ShowLtsConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLtsConfigsRequest struct{}"
	}

	return strings.Join([]string{"ShowLtsConfigsRequest", string(data)}, " ")
}
