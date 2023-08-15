package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterSettingRequest Request Object
type UpdateClusterSettingRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

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
