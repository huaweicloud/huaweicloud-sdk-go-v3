package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowResourceBindEnterpriseProjectRequest struct {
	// 企业项目ID

	EnterpriseProjectId string `json:"enterprise_project_id"`

	Body *ResqEpResouce `json:"body,omitempty"`
}

func (o ShowResourceBindEnterpriseProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceBindEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceBindEnterpriseProjectRequest", string(data)}, " ")
}
