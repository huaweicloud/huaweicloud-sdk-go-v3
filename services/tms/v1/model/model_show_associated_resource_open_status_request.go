package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssociatedResourceOpenStatusRequest Request Object
type ShowAssociatedResourceOpenStatusRequest struct {
}

func (o ShowAssociatedResourceOpenStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssociatedResourceOpenStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowAssociatedResourceOpenStatusRequest", string(data)}, " ")
}
