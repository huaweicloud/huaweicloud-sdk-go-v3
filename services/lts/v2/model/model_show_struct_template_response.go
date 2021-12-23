package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowStructTemplateResponse struct {
	// 结构化字段

	DemoFields *[]StructFieldInfoReturn `json:"demoFields,omitempty"`
	// 关键词详细信息

	TagFields *[]TagFieldsInfo `json:"tagFields,omitempty"`
	// 示例日志

	DemoLog *string `json:"demoLog,omitempty"`
	// 测试

	DemoLabel *string `json:"demoLabel,omitempty"`
	// id

	Id *string `json:"id,omitempty"`
	// 日志组ID

	LogGroupId *string `json:"logGroupId,omitempty"`
	// 结构化方式

	Rule *ShowStructTemplateRule `json:"rule,omitempty"`
	// kafka信息

	ClusterInfo *ShowStructTemplateclusterInfo `json:"cluster_info,omitempty"`
	// 日志流ID

	LogStreamId *string `json:"logStreamId,omitempty"`
	// 项目ID

	ProjectId *string `json:"projectId,omitempty"`
	// 测试

	TemplateName *string `json:"templateName,omitempty"`
	// 为了兼容前台数据格式

	Regex          *string `json:"regex,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowStructTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStructTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowStructTemplateResponse", string(data)}, " ")
}
