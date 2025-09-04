package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TestReportCustomDetailVo 实际的数据类型：单个对象，集合 或 NULL
type TestReportCustomDetailVo struct {

	// 测试报告自定义模块Uri
	Uri *string `json:"uri,omitempty"`

	// 测试报告自定义模块名称
	Name *string `json:"name,omitempty"`

	// 是否显示(0:否，1:是)
	Display *int32 `json:"display,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 附件信息
	Attachments *[]AttachmentVo `json:"attachments,omitempty"`

	// 创建人ID
	Creator *string `json:"creator,omitempty"`

	// 修改人ID
	Updator *string `json:"updator,omitempty"`

	// 测试报告uri
	TestReportUri *string `json:"test_report_uri,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 创建时间戳
	CreateTimestamp *int64 `json:"create_timestamp,omitempty"`

	// 创建人名
	CreatorName *string `json:"creator_name,omitempty"`

	// 修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 修改时间戳
	UpdateTimestamp *int64 `json:"update_timestamp,omitempty"`

	// 修改人名
	UpdatorName *string `json:"updator_name,omitempty"`
}

func (o TestReportCustomDetailVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestReportCustomDetailVo struct{}"
	}

	return strings.Join([]string{"TestReportCustomDetailVo", string(data)}, " ")
}
