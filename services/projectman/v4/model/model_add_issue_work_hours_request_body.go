package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddIssueWorkHoursRequestBody struct {

	// 工时开始日期，年-月-日
	StartDate string `json:"start_date"`

	// 工时结束日期，年-月-日
	DueDate string `json:"due_date"`

	// 工时总数（若工时日期范围包含多天，单日工时将设为“工时总数/天数”）
	WorkHours float64 `json:"work_hours"`

	// 工时类型id（项目预设工时类型id及名称对照：21:研发设计，22:后端开发，23:前端开发(Web)，24:前端开发(小程序)，25:前端开发(App)， 26:测试验证，27:缺陷修复，28:UI设计，29:会议，30:公共事务，31:培训，32:研究，33:其它，34:调休请假）
	WorkHoursTypeId *int32 `json:"work_hours_type_id,omitempty"`
}

func (o AddIssueWorkHoursRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddIssueWorkHoursRequestBody struct{}"
	}

	return strings.Join([]string{"AddIssueWorkHoursRequestBody", string(data)}, " ")
}
