package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityDatasourceProtectionDiagnoseResultRequest Request Object
type ShowSecurityDatasourceProtectionDiagnoseResultRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`
}

func (o ShowSecurityDatasourceProtectionDiagnoseResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityDatasourceProtectionDiagnoseResultRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityDatasourceProtectionDiagnoseResultRequest", string(data)}, " ")
}
