package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptSecurityApplicationRequest Request Object
type AcceptSecurityApplicationRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 工单id
	Id string `json:"id"`

	Body *ReasonDto `json:"body,omitempty"`
}

func (o AcceptSecurityApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptSecurityApplicationRequest struct{}"
	}

	return strings.Join([]string{"AcceptSecurityApplicationRequest", string(data)}, " ")
}
