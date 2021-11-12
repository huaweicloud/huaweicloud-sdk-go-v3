package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListConnectionsRequest struct {
}

func (o ListConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListConnectionsRequest", string(data)}, " ")
}
