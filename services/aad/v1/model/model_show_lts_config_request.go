package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLtsConfigRequest Request Object
type ShowLtsConfigRequest struct {

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowLtsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLtsConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowLtsConfigRequest", string(data)}, " ")
}
