package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRouteTags tags
type CreateRouteTags struct {

	// version
	Version *string `json:"version,omitempty"`
}

func (o CreateRouteTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRouteTags struct{}"
	}

	return strings.Join([]string{"CreateRouteTags", string(data)}, " ")
}
