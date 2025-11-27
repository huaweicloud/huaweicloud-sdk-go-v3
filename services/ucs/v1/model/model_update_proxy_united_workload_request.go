package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProxyUnitedWorkloadRequest Request Object
type UpdateProxyUnitedWorkloadRequest struct {

	// 容器舰队id
	Clustergroupid string `json:"clustergroupid"`

	Body *MutiWorkload `json:"body,omitempty"`
}

func (o UpdateProxyUnitedWorkloadRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProxyUnitedWorkloadRequest struct{}"
	}

	return strings.Join([]string{"UpdateProxyUnitedWorkloadRequest", string(data)}, " ")
}
