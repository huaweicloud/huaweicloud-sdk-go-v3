package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeProtectEipRequest struct {
	Body *EipOperateProtectReq `json:"body,omitempty"`
}

func (o ChangeProtectEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeProtectEipRequest struct{}"
	}

	return strings.Join([]string{"ChangeProtectEipRequest", string(data)}, " ")
}
