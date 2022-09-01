package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentDetailExtInfoV2 struct {

	// 国家码
	AreaCode *string `json:"area_code,omitempty" xml:"area_code"`

	// 提醒手机
	RemindMobile *string `json:"remind_mobile,omitempty" xml:"remind_mobile"`

	// 提醒邮箱
	RemindMail *string `json:"remind_mail,omitempty" xml:"remind_mail"`

	// 联系方式类型
	ContactType *string `json:"contact_type,omitempty" xml:"contact_type"`

	// 提醒时间
	RemindTime *string `json:"remind_time,omitempty" xml:"remind_time"`

	// 抄送邮箱
	CcEmail *string `json:"cc_email,omitempty" xml:"cc_email"`

	// ISV商品id
	CommodityId *string `json:"commodity_id,omitempty" xml:"commodity_id"`
}

func (o IncidentDetailExtInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentDetailExtInfoV2 struct{}"
	}

	return strings.Join([]string{"IncidentDetailExtInfoV2", string(data)}, " ")
}
