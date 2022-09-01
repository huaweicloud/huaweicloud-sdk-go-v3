package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteValueListRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 引用表id，通过查询引用表列表（ListValueList）接口获取
	Valuelistid string `json:"valuelistid" xml:"valuelistid"`
}

func (o DeleteValueListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteValueListRequest struct{}"
	}

	return strings.Join([]string{"DeleteValueListRequest", string(data)}, " ")
}
