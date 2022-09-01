package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowStructTemplateResponse struct {

	// 结构化字段
	DemoFields *[]StructFieldInfoReturn `json:"demoFields,omitempty" xml:"demoFields"`

	// 关键词详细信息
	TagFields *[]TagFieldsInfo `json:"tagFields,omitempty" xml:"tagFields"`

	// 示例日志
	DemoLog *string `json:"demoLog,omitempty" xml:"demoLog"`

	// 测试
	DemoLabel *string `json:"demoLabel,omitempty" xml:"demoLabel"`

	// id
	Id *string `json:"id,omitempty" xml:"id"`

	// 日志组ID
	LogGroupId *string `json:"logGroupId,omitempty" xml:"logGroupId"`

	Rule *ShowStructTemplateRule `json:"rule,omitempty" xml:"rule"`

	ClusterInfo *ShowStructTemplateclusterInfo `json:"cluster_info,omitempty" xml:"cluster_info"`

	// 日志流ID
	LogStreamId *string `json:"logStreamId,omitempty" xml:"logStreamId"`

	// 项目ID
	ProjectId *string `json:"projectId,omitempty" xml:"projectId"`

	// 测试
	TemplateName *string `json:"templateName,omitempty" xml:"templateName"`

	// 为了兼容前台数据格式
	Regex          *string `json:"regex,omitempty" xml:"regex"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowStructTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStructTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowStructTemplateResponse", string(data)}, " ")
}
