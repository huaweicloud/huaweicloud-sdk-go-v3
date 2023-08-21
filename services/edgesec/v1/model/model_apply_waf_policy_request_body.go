package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplyWafPolicyRequestBody struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 待关联域名的id
	Hosts []string `json:"hosts"`
}

func (o ApplyWafPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyWafPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"ApplyWafPolicyRequestBody", string(data)}, " ")
}
