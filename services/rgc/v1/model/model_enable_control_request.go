package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableControlRequest Request Object
type EnableControlRequest struct {
	Body *ControlOperateReqBody `json:"body,omitempty"`
}

func (o EnableControlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableControlRequest struct{}"
	}

	return strings.Join([]string{"EnableControlRequest", string(data)}, " ")
}
