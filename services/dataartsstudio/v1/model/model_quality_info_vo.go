package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QualityInfoVo struct {

	// 编码
	Id *int64 `json:"id,omitempty"`

	// 表id
	TableId *int64 `json:"table_id,omitempty"`

	// 属性id
	AttrId *int64 `json:"attr_id,omitempty"`

	BizType *BizTypeEnum `json:"biz_type,omitempty"`

	// 质量id
	DataQualityId int64 `json:"data_quality_id"`

	// 是否要显示  正则表达式
	ShowControl *int32 `json:"show_control,omitempty"`

	// 质量名称
	DataQualityName *string `json:"data_quality_name,omitempty"`

	// 告警配置
	AlertConf *string `json:"alert_conf,omitempty"`

	// 正则相关校验规则中正则配置
	Expression *string `json:"expression,omitempty"`

	// 扩展信息
	ExtendInfo *string `json:"extend_info,omitempty"`

	// 是否来源于数据标准质量配置
	FromStandard *bool `json:"from_standard,omitempty"`

	// 结果说明
	ResultDescription *string `json:"result_description,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人
	UpdateBy *string `json:"update_by,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o QualityInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QualityInfoVo struct{}"
	}

	return strings.Join([]string{"QualityInfoVo", string(data)}, " ")
}
