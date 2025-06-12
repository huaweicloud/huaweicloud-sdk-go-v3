package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobTotalRequest Request Object
type ShowJobTotalRequest struct {

	// 构建工程项目ID，36位数字、小写字母组合。
	BuildProjectId string `json:"build_project_id"`

	// 区间开始时间，格式yyyy-MM-dd HH:mm:ss。
	FromDate *string `json:"from_date,omitempty"`

	// 区间结束时间，格式yyyy-MM-dd HH:mm:ss。
	ToDate *string `json:"to_date,omitempty"`
}

func (o ShowJobTotalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobTotalRequest struct{}"
	}

	return strings.Join([]string{"ShowJobTotalRequest", string(data)}, " ")
}
