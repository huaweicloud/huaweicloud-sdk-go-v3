package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEpConfigsRequest Request Object
type ShowEpConfigsRequest struct {
}

func (o ShowEpConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEpConfigsRequest struct{}"
	}

	return strings.Join([]string{"ShowEpConfigsRequest", string(data)}, " ")
}
