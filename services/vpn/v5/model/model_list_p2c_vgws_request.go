package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListP2cVgwsRequest Request Object
type ListP2cVgwsRequest struct {
}

func (o ListP2cVgwsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListP2cVgwsRequest struct{}"
	}

	return strings.Join([]string{"ListP2cVgwsRequest", string(data)}, " ")
}
