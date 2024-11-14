package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyConfigurationToInstancesRequest Request Object
type ApplyConfigurationToInstancesRequest struct {

	// 参数模板ID。
	ConfigId string `json:"config_id"`

	Body *ApplyConfigurationToInstancesRequestBody `json:"body,omitempty"`
}

func (o ApplyConfigurationToInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyConfigurationToInstancesRequest struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationToInstancesRequest", string(data)}, " ")
}
