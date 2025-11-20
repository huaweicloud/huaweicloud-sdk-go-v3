package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteTransferAssetActionRequest Request Object
type ExecuteTransferAssetActionRequest struct {

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 控制。 cancel：取消资产转移，仅转移发起方可调用。 accept：接受资产转移，仅转移接受方可调用。 accept_confirm：确认接受资产转移，仅转移接受方可调用，仅需要计费的转移需再次确认。 reject: 拒绝资产转移，仅转移接受方可调用。
	Action string `json:"action"`

	Body *TransJobRejectBody `json:"body,omitempty"`
}

func (o ExecuteTransferAssetActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTransferAssetActionRequest struct{}"
	}

	return strings.Join([]string{"ExecuteTransferAssetActionRequest", string(data)}, " ")
}
