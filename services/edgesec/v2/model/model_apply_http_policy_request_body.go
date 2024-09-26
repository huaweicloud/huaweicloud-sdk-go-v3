package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplyHttpPolicyRequestBody struct {

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 待关联域名的id
	Hosts []string `json:"hosts"`
}

func (o ApplyHttpPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyHttpPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"ApplyHttpPolicyRequestBody", string(data)}, " ")
}
