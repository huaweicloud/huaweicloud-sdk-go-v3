package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceRequest Request Object
type ListInstanceRequest struct {
}

func (o ListInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceRequest", string(data)}, " ")
}
