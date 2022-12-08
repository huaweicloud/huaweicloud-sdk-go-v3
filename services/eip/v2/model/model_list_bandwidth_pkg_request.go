package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBandwidthPkgRequest struct {
}

func (o ListBandwidthPkgRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthPkgRequest struct{}"
	}

	return strings.Join([]string{"ListBandwidthPkgRequest", string(data)}, " ")
}
