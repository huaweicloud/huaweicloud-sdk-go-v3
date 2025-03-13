package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnpublishSecurityApplicationRequest Request Object
type UnpublishSecurityApplicationRequest struct {

	// 申请工单id
	Id string `json:"id"`

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`
}

func (o UnpublishSecurityApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnpublishSecurityApplicationRequest struct{}"
	}

	return strings.Join([]string{"UnpublishSecurityApplicationRequest", string(data)}, " ")
}
