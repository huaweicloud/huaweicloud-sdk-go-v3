package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReportBean struct {

	// 数据库ID
	DbIds *string `json:"db_ids,omitempty"`

	// 数据库名称
	DbNames *string `json:"db_names,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 完成时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 格式  - pdf: PDF文件  - zip: zip文件
	Format *string `json:"format,omitempty"`

	// 报表ID
	Id *string `json:"id,omitempty"`

	// 报表名称
	Name *string `json:"name,omitempty"`

	// 进度
	Percentum *int32 `json:"percentum,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 报表类型 - PDF: pdf - ZIP: zip
	TemplateType *string `json:"template_type,omitempty"`

	// 周期 - AUDIT_REPORT_DAY: 按天 - AUDIT_REPORT_WEEK： 按周 - AUDIT_REPORT_MONTH： 按月 - AUDIT_REPORT_YEAR：按年 - AUDIT_REPORT_REAL_TIME：立即
	Type *string `json:"type,omitempty"`

	// 地址URL
	Url *string `json:"url,omitempty"`
}

func (o ReportBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportBean struct{}"
	}

	return strings.Join([]string{"ReportBean", string(data)}, " ")
}
