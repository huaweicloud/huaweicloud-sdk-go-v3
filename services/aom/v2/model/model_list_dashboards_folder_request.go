package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDashboardsFolderRequest Request Object
type ListDashboardsFolderRequest struct {

	// 企业项目id。获取方式请参见：[获取企业项目ID](aom_04_0024.xml)。 - 查询单个企业项目下实例，填写企业项目id。 - 查询所有企业项目下实例，填写“all_granted_eps”。
	EnterpriseProjectId *string `json:"Enterprise-Project-Id,omitempty"`
}

func (o ListDashboardsFolderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDashboardsFolderRequest struct{}"
	}

	return strings.Join([]string{"ListDashboardsFolderRequest", string(data)}, " ")
}
