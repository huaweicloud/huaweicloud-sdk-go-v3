package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCustomRegexSetRequest Request Object
type BatchDeleteCustomRegexSetRequest struct {
	Body *AiguardCustomRegexSetBatchDeleteBody `json:"body,omitempty"`
}

func (o BatchDeleteCustomRegexSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCustomRegexSetRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteCustomRegexSetRequest", string(data)}, " ")
}
