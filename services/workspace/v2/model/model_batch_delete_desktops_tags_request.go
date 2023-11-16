package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDesktopsTagsRequest Request Object
type BatchDeleteDesktopsTagsRequest struct {
	Body *BatchDeleteDesktopsTagsReq `json:"body,omitempty"`
}

func (o BatchDeleteDesktopsTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDesktopsTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteDesktopsTagsRequest", string(data)}, " ")
}
