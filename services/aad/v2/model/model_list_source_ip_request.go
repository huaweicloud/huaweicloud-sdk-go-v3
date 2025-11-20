package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSourceIpRequest Request Object
type ListSourceIpRequest struct {
}

func (o ListSourceIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSourceIpRequest struct{}"
	}

	return strings.Join([]string{"ListSourceIpRequest", string(data)}, " ")
}
