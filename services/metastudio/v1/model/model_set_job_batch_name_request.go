package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetJobBatchNameRequest Request Object
type SetJobBatchNameRequest struct {

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	Body *SetJobBatchNameReq `json:"body,omitempty"`
}

func (o SetJobBatchNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetJobBatchNameRequest struct{}"
	}

	return strings.Join([]string{"SetJobBatchNameRequest", string(data)}, " ")
}
