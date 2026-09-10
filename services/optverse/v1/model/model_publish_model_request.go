package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishModelRequest Request Object
type PublishModelRequest struct {
	Body *PublishModelReq `json:"body,omitempty"`
}

func (o PublishModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishModelRequest struct{}"
	}

	return strings.Join([]string{"PublishModelRequest", string(data)}, " ")
}
