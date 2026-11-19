package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyGroupDetailResponse Response Object
type ListPolicyGroupDetailResponse struct {

	// **参数解释**： 策略组描述 **取值范围**： 字符长度1-64位
	Description *string `json:"description,omitempty"`

	// **参数解释**： 策略组ID **取值范围**： 字符长度1-128位
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**： 策略组名称 **取值范围**： 字符长度1-256位
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 支持的版本 **取值范围**： 字符长度1-64位
	SupportVersion *string `json:"support_version,omitempty"`

	// **参数解释**： 支持的操作系统 **取值范围**： 字符长度1-64位
	SupportOs *string `json:"support_os,omitempty"`

	// **参数解释**： 策略详情列表 **取值范围**： 不涉及
	DataList       *[]ShowPolicyDetailResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListPolicyGroupDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyGroupDetailResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyGroupDetailResponse", string(data)}, " ")
}
