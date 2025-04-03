package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowValueListRequest Request Object
type ShowValueListRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 引用表id，通过查询引用表列表（ListValueList）接口获取
	Valuelistid string `json:"valuelistid"`
}

func (o ShowValueListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowValueListRequest struct{}"
	}

	return strings.Join([]string{"ShowValueListRequest", string(data)}, " ")
}
