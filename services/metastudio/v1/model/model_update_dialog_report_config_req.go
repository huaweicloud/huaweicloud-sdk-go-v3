package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDialogReportConfigReq 修改对话结果上报配置请求。
type UpdateDialogReportConfigReq struct {

	// **参数解释**： 接收对话结果上报的obs桶名。 **约束限制**： 不涉及 **取值范围**： 字符长度1-64 **默认取值**： 不涉及
	ObsBucketName *string `json:"obs_bucket_name,omitempty"`

	// **参数解释**： 接收对话结果上报的obs终端节点。 **约束限制**： 需要为obs合法终端节点。 **取值范围**： 字符长度1-64 **默认取值**： 不涉及
	ObsEndpoint *string `json:"obs_endpoint,omitempty"`

	// 对话结果上报开关
	EnableDialogReport *bool `json:"enable_dialog_report,omitempty"`
}

func (o UpdateDialogReportConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDialogReportConfigReq struct{}"
	}

	return strings.Join([]string{"UpdateDialogReportConfigReq", string(data)}, " ")
}
