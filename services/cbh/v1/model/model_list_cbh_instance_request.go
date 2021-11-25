package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCbhInstanceRequest struct {
}

func (o ListCbhInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCbhInstanceRequest struct{}"
	}

	return strings.Join([]string{"ListCbhInstanceRequest", string(data)}, " ")
}
