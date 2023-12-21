package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebootInstanceRequest Request Object
type RebootInstanceRequest struct {
	Body *RebootCbhRequestBody `json:"body,omitempty"`
}

func (o RebootInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootInstanceRequest struct{}"
	}

	return strings.Join([]string{"RebootInstanceRequest", string(data)}, " ")
}
