package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateOrderIncidentV2Req struct {

	// 工单子类
	IncidentSubTypeId *string `json:"incident_sub_type_id,omitempty" xml:"incident_sub_type_id"`

	// 工单产品类型
	ProductCategoryId *string `json:"product_category_id,omitempty" xml:"product_category_id"`

	// 工单问题类型
	BusinessTypeId string `json:"business_type_id" xml:"business_type_id"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty" xml:"region_id"`

	// 问题描述
	SimpleDescription string `json:"simple_description" xml:"simple_description"`

	// 工单来源，当前固定为83aeb0f2834c4df49826c781d32a963e
	SourceId string `json:"source_id" xml:"source_id"`

	// 是否授权
	IsAuthorized *int32 `json:"is_authorized,omitempty" xml:"is_authorized"`

	// 机密信息内容
	AuthorizationContent *string `json:"authorization_content,omitempty" xml:"authorization_content"`

	// 提醒手机号
	RemindMobile *string `json:"remind_mobile,omitempty" xml:"remind_mobile"`

	// 提醒邮箱
	RemindMail *string `json:"remind_mail,omitempty" xml:"remind_mail"`

	// 提醒时间，如果是任意时间传0，如果是指定时间，传客户首选项对应时区的时间，比如09:00-18:00
	RemindTime *string `json:"remind_time,omitempty" xml:"remind_time"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 附件id列表
	AccessoryIds *[]string `json:"accessory_ids,omitempty" xml:"accessory_ids"`

	// 附加参数
	ExtendsMap map[string]interface{} `json:"extends_map,omitempty" xml:"extends_map"`

	// 扩展参数
	ExtensionMap map[string]interface{} `json:"extension_map,omitempty" xml:"extension_map"`

	// 严重性id
	SeverityId *string `json:"severity_id,omitempty" xml:"severity_id"`

	// 验证码
	VerifyCode *string `json:"verify_code,omitempty" xml:"verify_code"`

	// 国家码
	AreaCode *string `json:"area_code,omitempty" xml:"area_code"`
}

func (o CreateOrderIncidentV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrderIncidentV2Req struct{}"
	}

	return strings.Join([]string{"CreateOrderIncidentV2Req", string(data)}, " ")
}
