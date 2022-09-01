package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SetOffSiteBackupPolicyRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *SetOffSiteBackupPolicyRequestBody `json:"body,omitempty" xml:"body"`
}

func (o SetOffSiteBackupPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetOffSiteBackupPolicyRequest struct{}"
	}

	return strings.Join([]string{"SetOffSiteBackupPolicyRequest", string(data)}, " ")
}
