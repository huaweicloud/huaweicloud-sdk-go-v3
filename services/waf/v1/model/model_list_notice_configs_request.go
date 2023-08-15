package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNoticeConfigsRequest Request Object
type ListNoticeConfigsRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListNoticeConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNoticeConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListNoticeConfigsRequest", string(data)}, " ")
}
