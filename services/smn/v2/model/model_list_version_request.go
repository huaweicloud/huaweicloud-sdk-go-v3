package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVersionRequest Request Object
type ListVersionRequest struct {
}

func (o ListVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVersionRequest struct{}"
	}

	return strings.Join([]string{"ListVersionRequest", string(data)}, " ")
}
