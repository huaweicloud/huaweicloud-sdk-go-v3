package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateOrderIncidentV2Req struct {

	// 工单子类，通过\"查询工单类目列表\"接口获取
	IncidentSubTypeId *string `json:"incident_sub_type_id,omitempty"`

	// 工单产品类型，通过\"查询工单类目列表\"接口获取
	ProductCategoryId *string `json:"product_category_id,omitempty"`

	// 工单问题类型，通过\"查询问题类型列表\"接口获取
	BusinessTypeId string `json:"business_type_id"`

	// 区域ID，根据business_type_id对应工单类型是5则必填，通过\"查询问题类型列表\"查看
	RegionId *string `json:"region_id,omitempty"`

	// 问题描述
	SimpleDescription string `json:"simple_description"`

	// 工单来源，当前固定为83aeb0f2834c4df49826c781d32a963e
	SourceId string `json:"source_id"`

	// 是否授权
	IsAuthorized *int32 `json:"is_authorized,omitempty"`

	// 机密信息内容
	AuthorizationContent *string `json:"authorization_content,omitempty"`

	// 提醒手机号
	RemindMobile *string `json:"remind_mobile,omitempty"`

	// 提醒邮箱
	RemindMail *string `json:"remind_mail,omitempty"`

	// 提醒时间，如果是任意时间传0，如果是指定时间，传客户首选项对应时区的时间，比如09:00-18:00
	RemindTime *string `json:"remind_time,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 附件id列表，\"上传附件\"接口返回的id
	AccessoryIds *[]string `json:"accessory_ids,omitempty"`

	// 附加参数
	ExtendsMap map[string]interface{} `json:"extends_map,omitempty"`

	// 扩展参数
	ExtensionMap map[string]interface{} `json:"extension_map,omitempty"`

	// 严重性id，通过\"查询问题严重性列表\"接口获取
	SeverityId *string `json:"severity_id,omitempty"`

	// 验证码，如果是非注册联系方式，需要通过\"获取验证码\"获取验证码
	VerifyCode *string `json:"verify_code,omitempty"`

	// 国家码
	AreaCode *string `json:"area_code,omitempty"`
}

func (o CreateOrderIncidentV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrderIncidentV2Req struct{}"
	}

	return strings.Join([]string{"CreateOrderIncidentV2Req", string(data)}, " ")
}
