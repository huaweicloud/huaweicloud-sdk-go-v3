package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSourceIpsRequest Request Object
type ListSourceIpsRequest struct {
}

func (o ListSourceIpsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSourceIpsRequest struct{}"
	}

	return strings.Join([]string{"ListSourceIpsRequest", string(data)}, " ")
}
