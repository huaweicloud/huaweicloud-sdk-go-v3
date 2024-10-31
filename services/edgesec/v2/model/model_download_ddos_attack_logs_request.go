package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadDdosAttackLogsRequest Request Object
type DownloadDdosAttackLogsRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id，默认为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 起始时间(13位时间戳)
	StartTime int64 `json:"start_time"`

	// 结束时间(13位时间戳)
	EndTime int64 `json:"end_time"`

	// 查询列表的偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 查询列表每一页的条数
	Limit *int32 `json:"limit,omitempty"`
}

func (o DownloadDdosAttackLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadDdosAttackLogsRequest struct{}"
	}

	return strings.Join([]string{"DownloadDdosAttackLogsRequest", string(data)}, " ")
}
