package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TestReportCustomDetailInfo struct {

	// 测试报告自定义模块名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 附件信息
	Attachments *[]AttachmentInfo `json:"attachments,omitempty"`
}

func (o TestReportCustomDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestReportCustomDetailInfo struct{}"
	}

	return strings.Join([]string{"TestReportCustomDetailInfo", string(data)}, " ")
}
