package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 输出为hosting类型时的配置和展示信息
type TaskOutputHostingForDisplay struct {

	// 作业所有结果文件所在的OBS桶和路径
	Obs *[]TaskOutputHostingForDisplayObs `json:"obs,omitempty"`

	// 作业结果文件的过期时间
	ResultJsonOverdueAt *int64 `json:"result_json_overdue_at,omitempty"`

	// 作业输出数据类别的列表，当输出类型下有这个列表时，表示希望这个输出类型下存放dataCategory列表内的数据，部分服务需要
	DataCategory *[]string `json:"data_category,omitempty"`
}

func (o TaskOutputHostingForDisplay) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskOutputHostingForDisplay struct{}"
	}

	return strings.Join([]string{"TaskOutputHostingForDisplay", string(data)}, " ")
}
