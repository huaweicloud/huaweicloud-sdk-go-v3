package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDashboardsFolderRequest Request Object
type DeleteDashboardsFolderRequest struct {

	// 仪表盘分组id。
	FolderId string `json:"folder_id"`

	// 企业项目id。获取方式请参见：[获取企业项目ID](aom_04_0024.xml)。  - 删除单个企业项目下实例，填写企业项目id。  - 不填时，默认删除企业项目id为0的企业项目下实例。
	EnterpriseProjectId *string `json:"Enterprise-Project-Id,omitempty"`

	// 是否删除仪表盘分组下的仪表盘。
	DeleteAll bool `json:"delete_all"`
}

func (o DeleteDashboardsFolderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDashboardsFolderRequest struct{}"
	}

	return strings.Join([]string{"DeleteDashboardsFolderRequest", string(data)}, " ")
}
