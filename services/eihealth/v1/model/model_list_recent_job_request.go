package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRecentJobRequest struct {

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListRecentJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecentJobRequest struct{}"
	}

	return strings.Join([]string{"ListRecentJobRequest", string(data)}, " ")
}
