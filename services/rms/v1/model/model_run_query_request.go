package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunQueryRequest struct {
	Body *QueryRunRequestBody `json:"body,omitempty"`
}

func (o RunQueryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunQueryRequest struct{}"
	}

	return strings.Join([]string{"RunQueryRequest", string(data)}, " ")
}
