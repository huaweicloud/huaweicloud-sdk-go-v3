package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FindingDetails 访问分析结果的详细信息。
type FindingDetails struct {
	ExternalAccessDetails *ExternalAccessDetails `json:"external_access_details,omitempty"`

	UnusedIamUserAccessKeyDetails *UnusedIamUserAccessKeyDetails `json:"unused_iam_user_access_key_details,omitempty"`

	UnusedIamUserPasswordDetails *UnusedIamUserPasswordDetails `json:"unused_iam_user_password_details,omitempty"`
}

func (o FindingDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FindingDetails struct{}"
	}

	return strings.Join([]string{"FindingDetails", string(data)}, " ")
}
