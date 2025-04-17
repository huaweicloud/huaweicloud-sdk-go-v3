package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogConfigRequest Request Object
type UpdateLogConfigRequest struct {

	// 企业项目id
	EnterpriseProjectId string `json:"enterprise_project_id"`

	Body *LtsConfigRequestAndResponse `json:"body,omitempty"`
}

func (o UpdateLogConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateLogConfigRequest", string(data)}, " ")
}
