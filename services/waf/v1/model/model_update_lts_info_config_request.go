package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateLtsInfoConfigRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// lts配置信息id，通过ShowLtsInfoConfig获取
	LtsconfigId string `json:"ltsconfig_id"`

	Body *UpdateLtsInfoConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateLtsInfoConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLtsInfoConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateLtsInfoConfigRequest", string(data)}, " ")
}
