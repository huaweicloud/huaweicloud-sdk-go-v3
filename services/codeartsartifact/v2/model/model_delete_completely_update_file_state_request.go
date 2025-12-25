package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCompletelyUpdateFileStateRequest Request Object
type DeleteCompletelyUpdateFileStateRequest struct {
	Body *[]string `json:"body,omitempty"`
}

func (o DeleteCompletelyUpdateFileStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCompletelyUpdateFileStateRequest struct{}"
	}

	return strings.Join([]string{"DeleteCompletelyUpdateFileStateRequest", string(data)}, " ")
}
