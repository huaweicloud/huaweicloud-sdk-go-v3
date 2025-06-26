package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFeatureGatesRequest Request Object
type ListFeatureGatesRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`
}

func (o ListFeatureGatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFeatureGatesRequest struct{}"
	}

	return strings.Join([]string{"ListFeatureGatesRequest", string(data)}, " ")
}
