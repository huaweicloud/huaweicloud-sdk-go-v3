package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityCertification 安全认证诊断结果
type SecurityCertification struct {
	Result *DiagnoseResult `json:"result,omitempty"`

	// kerberos信息
	KerberosInfo *[]KerberosStatus `json:"kerberos_info,omitempty"`
}

func (o SecurityCertification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCertification struct{}"
	}

	return strings.Join([]string{"SecurityCertification", string(data)}, " ")
}
