package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDiskTypeRequest Request Object
type ListDiskTypeRequest struct {
}

func (o ListDiskTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiskTypeRequest struct{}"
	}

	return strings.Join([]string{"ListDiskTypeRequest", string(data)}, " ")
}
