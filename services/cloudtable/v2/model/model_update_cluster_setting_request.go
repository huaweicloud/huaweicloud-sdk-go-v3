package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateClusterSettingRequest struct {

	// 租户ID
	ProjectId string `json:"projectId"`

	// 集群ID
	ClusterId string `json:"clusterId"`

	// 语言类型
	XLanguage *string `json:"X-Language,omitempty"`

	Body *HbaseModifySettingV2Req `json:"body,omitempty"`
}

func (o UpdateClusterSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterSettingRequest struct{}"
	}

	return strings.Join([]string{"UpdateClusterSettingRequest", string(data)}, " ")
}
