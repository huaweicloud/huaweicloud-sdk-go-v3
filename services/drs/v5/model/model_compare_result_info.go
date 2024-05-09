package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareResultInfo 查询对比任务信息体。
type CompareResultInfo struct {
	ObjectsCompareOverviewInfo *ObjectsCompareTaskInfo `json:"objects_compare_overview_info,omitempty"`

	// 对象级对比详情信息体。
	ObjectsCompareDetailInfos *[]ObjectsCompareDetailInfo `json:"objects_compare_detail_infos,omitempty"`

	// 数据对比任务列表。
	DataCompareTaskList *[]CompareJobInfo `json:"data_compare_task_list,omitempty"`

	// 行数对比概览信息体。
	LineCompareOverviewInfos *[]LineCompareOverviewInfo `json:"line_compare_overview_infos,omitempty"`

	// 行数对比任务表级详情。
	LineCompareDetailInfos *[]TableLineCompareDetailInfo `json:"line_compare_detail_infos,omitempty"`

	// 内容对比概览信息体。
	ContentCompareOverviewInfos *[]ContentCompareOverviewInfo `json:"content_compare_overview_infos,omitempty"`

	// 内容对比详情信息体。
	ContentCompareDetailInfos *[]ContentCompareDetailInfo `json:"content_compare_detail_infos,omitempty"`

	ContentDiffDetailInfo *ContentDiffDetailInfo `json:"content_diff_detail_info,omitempty"`
}

func (o CompareResultInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareResultInfo struct{}"
	}

	return strings.Join([]string{"CompareResultInfo", string(data)}, " ")
}
