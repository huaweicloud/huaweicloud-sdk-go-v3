package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteBlackWhiteListsRequest Request Object
type BatchDeleteBlackWhiteListsRequest struct {
	Body *DeleteBlackWhiteDto `json:"body,omitempty"`
}

func (o BatchDeleteBlackWhiteListsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBlackWhiteListsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteBlackWhiteListsRequest", string(data)}, " ")
}
