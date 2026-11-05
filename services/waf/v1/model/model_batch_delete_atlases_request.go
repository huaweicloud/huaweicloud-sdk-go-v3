package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAtlasesRequest Request Object
type BatchDeleteAtlasesRequest struct {
	Body *AiguardAtlasBatchDeleteBody `json:"body,omitempty"`
}

func (o BatchDeleteAtlasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAtlasesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteAtlasesRequest", string(data)}, " ")
}
