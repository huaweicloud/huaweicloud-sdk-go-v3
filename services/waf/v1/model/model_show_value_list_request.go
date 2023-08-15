package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowValueListRequest Request Object
type ShowValueListRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
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
