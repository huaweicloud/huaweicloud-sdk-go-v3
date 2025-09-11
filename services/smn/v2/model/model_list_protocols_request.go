package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProtocolsRequest Request Object
type ListProtocolsRequest struct {
}

func (o ListProtocolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtocolsRequest struct{}"
	}

	return strings.Join([]string{"ListProtocolsRequest", string(data)}, " ")
}
