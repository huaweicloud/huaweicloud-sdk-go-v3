package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupportedMetricNamesRequest Request Object
type ListSupportedMetricNamesRequest struct {
}

func (o ListSupportedMetricNamesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportedMetricNamesRequest struct{}"
	}

	return strings.Join([]string{"ListSupportedMetricNamesRequest", string(data)}, " ")
}
