package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProxyUnitedWorkloadRequest Request Object
type CreateProxyUnitedWorkloadRequest struct {

	// 容器舰队id
	Clustergroupid string `json:"clustergroupid"`

	Body *MutiWorkload `json:"body,omitempty"`
}

func (o CreateProxyUnitedWorkloadRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProxyUnitedWorkloadRequest struct{}"
	}

	return strings.Join([]string{"CreateProxyUnitedWorkloadRequest", string(data)}, " ")
}
