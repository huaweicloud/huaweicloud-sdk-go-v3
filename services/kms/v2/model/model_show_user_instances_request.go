package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowUserInstancesRequest struct {
}

func (o ShowUserInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserInstancesRequest struct{}"
	}

	return strings.Join([]string{"ShowUserInstancesRequest", string(data)}, " ")
}
