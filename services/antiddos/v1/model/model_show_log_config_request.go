package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogConfigRequest Request Object
type ShowLogConfigRequest struct {

	// 企业项目id
	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o ShowLogConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowLogConfigRequest", string(data)}, " ")
}
