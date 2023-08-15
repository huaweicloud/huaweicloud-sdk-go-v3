package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOriginHostRequest Request Object
type UpdateOriginHostRequest struct {

	// 加速域名ID。
	DomainId string `json:"domain_id"`

	// 当用户开启企业项目功能时，该参数生效，表示修改当前企业项目下加速域名的配置，\"all\"代表所有项目。注意：当使用子帐号调用接口时，该参数必传。  您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *OriginHostRequest `json:"body,omitempty"`
}

func (o UpdateOriginHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOriginHostRequest struct{}"
	}

	return strings.Join([]string{"UpdateOriginHostRequest", string(data)}, " ")
}
