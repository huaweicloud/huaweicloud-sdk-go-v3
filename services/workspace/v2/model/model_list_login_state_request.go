package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLoginStateRequest Request Object
type ListLoginStateRequest struct {
}

func (o ListLoginStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoginStateRequest struct{}"
	}

	return strings.Join([]string{"ListLoginStateRequest", string(data)}, " ")
}
