package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowClusterSettingRequest struct {

	// 租户ID
	ProjectId string `json:"projectId"`

	// 项目ID
	ClusterId string `json:"clusterId"`

	// 语言类型
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowClusterSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterSettingRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterSettingRequest", string(data)}, " ")
}
