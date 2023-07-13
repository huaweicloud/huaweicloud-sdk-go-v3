package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogtankRequest Request Object
type UpdateLogtankRequest struct {

	// 云日志ID。
	LogtankId string `json:"logtank_id"`

	Body *UpdateLogtankRequestBody `json:"body,omitempty"`
}

func (o UpdateLogtankRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogtankRequest struct{}"
	}

	return strings.Join([]string{"UpdateLogtankRequest", string(data)}, " ")
}
