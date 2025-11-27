package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePolicyDefinationRequest Request Object
type UpdatePolicyDefinationRequest struct {

	// 策略定义id
	Policydefinitionid string `json:"policydefinitionid"`

	Body *UcsConstraintTemplateTagRequest `json:"body,omitempty"`
}

func (o UpdatePolicyDefinationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyDefinationRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyDefinationRequest", string(data)}, " ")
}
