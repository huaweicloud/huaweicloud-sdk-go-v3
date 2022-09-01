package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchSetAttributesRequest struct {
	Body *BatchSetAttributesReq `json:"body,omitempty" xml:"body"`
}

func (o BatchSetAttributesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetAttributesRequest struct{}"
	}

	return strings.Join([]string{"BatchSetAttributesRequest", string(data)}, " ")
}
