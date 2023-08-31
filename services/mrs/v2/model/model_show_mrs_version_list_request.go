package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMrsVersionListRequest Request Object
type ShowMrsVersionListRequest struct {
}

func (o ShowMrsVersionListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMrsVersionListRequest struct{}"
	}

	return strings.Join([]string{"ShowMrsVersionListRequest", string(data)}, " ")
}
