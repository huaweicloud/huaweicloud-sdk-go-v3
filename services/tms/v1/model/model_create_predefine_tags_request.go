package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePredefineTagsRequest struct {
	Body *ReqCreatePredefineTag `json:"body,omitempty"`
}

func (o CreatePredefineTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePredefineTagsRequest struct{}"
	}

	return strings.Join([]string{"CreatePredefineTagsRequest", string(data)}, " ")
}
