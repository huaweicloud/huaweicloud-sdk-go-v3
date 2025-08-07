package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEnterpriseProjectRequest Request Object
type DeleteEnterpriseProjectRequest struct {

	// 企业项目ID，不能为0。 可以通过查询企业项目列表接口获取。
	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o DeleteEnterpriseProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"DeleteEnterpriseProjectRequest", string(data)}, " ")
}
