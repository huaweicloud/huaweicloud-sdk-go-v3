package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDialogReportConfigResponse Response Object
type ShowDialogReportConfigResponse struct {

	// 租户ID。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 接收对话结果上报的obs桶名。 **约束限制**： 不涉及 **取值范围**： 字符长度1-64 **默认取值**： 不涉及
	ObsBucketName *string `json:"obs_bucket_name,omitempty"`

	// **参数解释**： 接收对话结果上报的obs终端节点。 **约束限制**： 需要为obs合法终端节点。 **取值范围**： 字符长度1-64 **默认取值**： 不涉及
	ObsEndpoint *string `json:"obs_endpoint,omitempty"`

	// 对话结果上报开关
	EnableDialogReport *bool `json:"enable_dialog_report,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDialogReportConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDialogReportConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowDialogReportConfigResponse", string(data)}, " ")
}
