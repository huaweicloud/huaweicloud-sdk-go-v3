package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateCaseContactInfoReq struct {

	// 国家码
	AreaCode *string `json:"area_code,omitempty"`

	// 联系手机号
	RemindMobile *string `json:"remind_mobile,omitempty"`

	// 联系邮箱
	RemindMail *string `json:"remind_mail,omitempty"`

	// 联系时间
	RemindTime *string `json:"remind_time,omitempty"`

	// 组id
	GroupId *string `json:"group_id,omitempty"`

	// 扩展参数
	ExtensionMap map[string]interface{} `json:"extension_map,omitempty"`
}

func (o UpdateCaseContactInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCaseContactInfoReq struct{}"
	}

	return strings.Join([]string{"UpdateCaseContactInfoReq", string(data)}, " ")
}
