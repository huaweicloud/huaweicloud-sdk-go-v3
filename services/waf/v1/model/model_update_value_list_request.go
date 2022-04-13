package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateValueListRequest struct {
	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 引用表id，通过查询引用表列表（ListValueList）接口获取

	Valuelistid string `json:"valuelistid"`

	Body *UpdateValueListRequestBody `json:"body,omitempty"`
}

func (o UpdateValueListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateValueListRequest struct{}"
	}

	return strings.Join([]string{"UpdateValueListRequest", string(data)}, " ")
}
