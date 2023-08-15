package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunIefJobActionCallBackRequest Request Object
type RunIefJobActionCallBackRequest struct {
	Body *IefFlinkJobMessagesReq `json:"body,omitempty"`
}

func (o RunIefJobActionCallBackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunIefJobActionCallBackRequest struct{}"
	}

	return strings.Join([]string{"RunIefJobActionCallBackRequest", string(data)}, " ")
}
