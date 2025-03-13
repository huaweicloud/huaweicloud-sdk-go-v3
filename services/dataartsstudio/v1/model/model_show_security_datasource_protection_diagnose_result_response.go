package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityDatasourceProtectionDiagnoseResultResponse Response Object
type ShowSecurityDatasourceProtectionDiagnoseResultResponse struct {

	// 诊断任务id
	TaskId *string `json:"task_id,omitempty"`

	// 是否正在扫描
	Scanning *bool `json:"scanning,omitempty"`

	// 最新检测时间。
	CheckTime *int64 `json:"check_time,omitempty"`

	Kerberos *SecurityCertification `json:"kerberos,omitempty"`

	PublicNetworkAccess *PublicNetworkAccess `json:"public_network_access,omitempty"`

	SecurityGroup  *SecurityGroupResult `json:"security_group,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowSecurityDatasourceProtectionDiagnoseResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityDatasourceProtectionDiagnoseResultResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityDatasourceProtectionDiagnoseResultResponse", string(data)}, " ")
}
