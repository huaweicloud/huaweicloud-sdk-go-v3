package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaListVersionsRequest Request Object
type NovaListVersionsRequest struct {
}

func (o NovaListVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaListVersionsRequest struct{}"
	}

	return strings.Join([]string{"NovaListVersionsRequest", string(data)}, " ")
}
