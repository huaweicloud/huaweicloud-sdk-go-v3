package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobRequest Request Object
type ListJobRequest struct {
}

func (o ListJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobRequest struct{}"
	}

	return strings.Join([]string{"ListJobRequest", string(data)}, " ")
}
