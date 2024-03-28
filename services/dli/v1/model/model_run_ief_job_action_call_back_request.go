package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunIefJobActionCallBackRequest Request Object
type RunIefJobActionCallBackRequest struct {
	Body *RunIefJobActionCallBackRequestBody `json:"body,omitempty"`
}

func (o RunIefJobActionCallBackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunIefJobActionCallBackRequest struct{}"
	}

	return strings.Join([]string{"RunIefJobActionCallBackRequest", string(data)}, " ")
}
