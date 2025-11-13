package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLtsConfigRequest Request Object
type UpdateLtsConfigRequest struct {

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *UpdateLtsConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateLtsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLtsConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateLtsConfigRequest", string(data)}, " ")
}
