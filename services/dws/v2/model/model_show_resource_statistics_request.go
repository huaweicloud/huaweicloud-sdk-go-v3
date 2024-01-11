package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceStatisticsRequest Request Object
type ShowResourceStatisticsRequest struct {

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`
}

func (o ShowResourceStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceStatisticsRequest", string(data)}, " ")
}
