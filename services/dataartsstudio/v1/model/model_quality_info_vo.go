package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QualityInfoVo struct {

	// 编码ID，填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 表ID，只读，填写String类型替代Long类型。
	TableId *string `json:"table_id,omitempty"`

	// 属性ID，只读，填写String类型替代Long类型。
	AttrId *string `json:"attr_id,omitempty"`

	BizType *BizTypeEnum `json:"biz_type,omitempty"`

	// 质量ID，填写String类型替代Long类型。
	DataQualityId string `json:"data_quality_id"`

	// 是否要显示正则表达式。
	ShowControl *int32 `json:"show_control,omitempty"`

	// 质量名称。
	DataQualityName *string `json:"data_quality_name,omitempty"`

	// 告警配置。
	AlertConf *string `json:"alert_conf,omitempty"`

	// 正则相关校验规则中正则配置。
	Expression *string `json:"expression,omitempty"`

	// 扩展信息。
	ExtendInfo *string `json:"extend_info,omitempty"`

	// 是否来源于数据标准质量配置，只读。
	FromStandard *bool `json:"from_standard,omitempty"`

	// 结果说明。
	ResultDescription *string `json:"result_description,omitempty"`

	// 创建人，只读。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人，只读。
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o QualityInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QualityInfoVo struct{}"
	}

	return strings.Join([]string{"QualityInfoVo", string(data)}, " ")
}
