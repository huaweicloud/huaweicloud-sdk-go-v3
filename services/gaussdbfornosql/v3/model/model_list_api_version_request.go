package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListApiVersionRequest struct {
}

func (o ListApiVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionRequest struct{}"
	}

	return strings.Join([]string{"ListApiVersionRequest", string(data)}, " ")
}
