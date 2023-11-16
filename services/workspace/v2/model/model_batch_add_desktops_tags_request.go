package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddDesktopsTagsRequest Request Object
type BatchAddDesktopsTagsRequest struct {
	Body *BatchAddDesktopsTagsReq `json:"body,omitempty"`
}

func (o BatchAddDesktopsTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddDesktopsTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchAddDesktopsTagsRequest", string(data)}, " ")
}
