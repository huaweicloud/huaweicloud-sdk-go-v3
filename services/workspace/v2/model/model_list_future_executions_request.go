package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFutureExecutionsRequest Request Object
type ListFutureExecutionsRequest struct {
	Body *ListFutureExecutionsReq `json:"body,omitempty"`
}

func (o ListFutureExecutionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFutureExecutionsRequest struct{}"
	}

	return strings.Join([]string{"ListFutureExecutionsRequest", string(data)}, " ")
}
