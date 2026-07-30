package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAssociatedResourceOpenStatusRequest Request Object
type ChangeAssociatedResourceOpenStatusRequest struct {
	Body *ReqAssociatedResourceOpenStatus `json:"body,omitempty"`
}

func (o ChangeAssociatedResourceOpenStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAssociatedResourceOpenStatusRequest struct{}"
	}

	return strings.Join([]string{"ChangeAssociatedResourceOpenStatusRequest", string(data)}, " ")
}
