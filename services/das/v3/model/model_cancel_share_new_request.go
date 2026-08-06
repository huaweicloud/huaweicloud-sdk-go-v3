package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelShareNewRequest Request Object
type CancelShareNewRequest struct {
	Body *CancelShareNewRequestBody `json:"body,omitempty"`
}

func (o CancelShareNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelShareNewRequest struct{}"
	}

	return strings.Join([]string{"CancelShareNewRequest", string(data)}, " ")
}
