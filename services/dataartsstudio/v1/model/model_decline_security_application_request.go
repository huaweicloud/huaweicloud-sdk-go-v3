package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeclineSecurityApplicationRequest Request Object
type DeclineSecurityApplicationRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 工单id
	Id string `json:"id"`

	Body *ReasonDto `json:"body,omitempty"`
}

func (o DeclineSecurityApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeclineSecurityApplicationRequest struct{}"
	}

	return strings.Join([]string{"DeclineSecurityApplicationRequest", string(data)}, " ")
}
