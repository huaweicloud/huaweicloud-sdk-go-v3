package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPipleineBuildResultRequest struct {
	// 语言类型 中文:zh-cn 英文:en-us，默认en-us

	XLanguage *string `json:"X-Language,omitempty"`
	// 项目id

	ProjectId string `json:"project_id"`
	// 起始日期,起始日期和结束日期间隔不超过一个月，查询包含起始日期

	StartDate string `json:"start_date"`
	// 结束日期，起始日期和结束日期间隔不超过一个月，查询包含结束日期

	EndDate string `json:"end_date"`
	// 偏移量,表示从此偏移量开始查询,offset大于等于0

	Offset int32 `json:"offset"`
	// 每次查询的条目数量

	Limit int32 `json:"limit"`
}

func (o ListPipleineBuildResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipleineBuildResultRequest struct{}"
	}

	return strings.Join([]string{"ListPipleineBuildResultRequest", string(data)}, " ")
}
