package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterSettingRequest Request Object
type ShowClusterSettingRequest struct {

	// 项目ID
	ClusterId string `json:"cluster_id"`

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
