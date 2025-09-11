package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableResourceRequest Request Object
type ShowAvailableResourceRequest struct {
}

func (o ShowAvailableResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableResourceRequest struct{}"
	}

	return strings.Join([]string{"ShowAvailableResourceRequest", string(data)}, " ")
}
