package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeSecurityGroupRequest Request Object
type ChangeSecurityGroupRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	Body *ChangeSecurityGroupRequestBody `json:"body,omitempty"`
}

func (o ChangeSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"ChangeSecurityGroupRequest", string(data)}, " ")
}
