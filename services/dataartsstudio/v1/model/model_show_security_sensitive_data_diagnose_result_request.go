package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecuritySensitiveDataDiagnoseResultRequest Request Object
type ShowSecuritySensitiveDataDiagnoseResultRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`
}

func (o ShowSecuritySensitiveDataDiagnoseResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecuritySensitiveDataDiagnoseResultRequest struct{}"
	}

	return strings.Join([]string{"ShowSecuritySensitiveDataDiagnoseResultRequest", string(data)}, " ")
}
