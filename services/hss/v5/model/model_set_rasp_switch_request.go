package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRaspSwitchRequest Request Object
type SetRaspSwitchRequest struct {

	// Region Id
	Region string `json:"region"`

	// 企业项目
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *SetRaspSwitchRequestInfo `json:"body,omitempty"`
}

func (o SetRaspSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRaspSwitchRequest struct{}"
	}

	return strings.Join([]string{"SetRaspSwitchRequest", string(data)}, " ")
}
