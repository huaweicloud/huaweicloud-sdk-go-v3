package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportStatsOpenRequest Request Object
type ExportStatsOpenRequest struct {

	// **参数解释：** 企业项目id > 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id  **约束限制：** - 当用户开启企业项目功能时，该参数生效，表示查询资源所属项目 - 当使用子账号调用接口时，该参数必传 **取值范围：** all表示所有项目 **默认取值：** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *ExportVo `json:"body,omitempty"`
}

func (o ExportStatsOpenRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportStatsOpenRequest struct{}"
	}

	return strings.Join([]string{"ExportStatsOpenRequest", string(data)}, " ")
}
