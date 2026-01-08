package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportAppUserAccessDataRequestBody 创建导出应用接入统计任务请求体。
type ExportAppUserAccessDataRequestBody struct {

	// 查询起始时间，格式为：UTC格式，例如\"2022-05-11T11:45:42Z\"。
	StartTime string `json:"start_time"`

	// 查询截至时间，格式为：UTC格式，例如\"2022-05-11T11:45:42Z\"。
	EndTime string `json:"end_time"`

	// 用户名(模糊匹配)。
	Username *string `json:"username,omitempty"`

	// 按照指标进行排序。 * `access_failed_count` -  按照接入失败数排序 * `last_access_failed_time` -  按照最近接入失败时间排序
	SortField *string `json:"sort_field,omitempty"`

	// 按照指标进行排序的方向;需配合sort_field一起使用。 * `DESC` - 降序返回数据 * `ASC` -  升序返回数据
	SortType *string `json:"sort_type,omitempty"`
}

func (o ExportAppUserAccessDataRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportAppUserAccessDataRequestBody struct{}"
	}

	return strings.Join([]string{"ExportAppUserAccessDataRequestBody", string(data)}, " ")
}
