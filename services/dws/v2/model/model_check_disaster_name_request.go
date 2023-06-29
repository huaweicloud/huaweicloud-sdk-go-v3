package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDisasterNameRequest Request Object
type CheckDisasterNameRequest struct {

	// 容灾名称
	DrName string `json:"dr_name"`

	// 容灾类型
	Type *string `json:"type,omitempty"`

	// 备集群所在region
	StandbyRegion *string `json:"standby_region,omitempty"`

	// 备集群所在项目ID
	StandbyProjectId *string `json:"standby_project_id,omitempty"`
}

func (o CheckDisasterNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDisasterNameRequest struct{}"
	}

	return strings.Join([]string{"CheckDisasterNameRequest", string(data)}, " ")
}
