package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStructTemplateResponse Response Object
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

	Rule *ShowStructTemplateRule `json:"rule,omitempty"`

	ClusterInfo *ShowStructTemplateclusterInfo `json:"cluster_info,omitempty"`

	// 日志流ID
	LogStreamId *string `json:"logStreamId,omitempty"`

	// 项目ID
	ProjectId *string `json:"projectId,omitempty"`

	// 测试
	TemplateName *string `json:"templateName,omitempty"`

	// 为了兼容前台数据格式
	Regex *string `json:"regex,omitempty"`

	CustomTimeInfo *CustomTimeInfo `json:"custom_time_info,omitempty"`

	// **参数解释：** 是否上传原始日志。 **取值范围：** - true - fasle
	UploadOriginalLog *bool `json:"uploadOriginalLog,omitempty"`

	// **参数解释：** 是否将解析失败的原始上传到指定系统字段日志。 **取值范围：** - true - fasle
	UploadParseFailedLog *bool `json:"uploadParseFailedLog,omitempty"`
	HttpStatusCode       int   `json:"-"`
}

func (o ShowStructTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStructTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowStructTemplateResponse", string(data)}, " ")
}
