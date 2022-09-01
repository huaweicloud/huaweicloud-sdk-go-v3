package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContactV2 struct {

	// 联系方式类型
	ContactWay *int32 `json:"contact_way,omitempty" xml:"contact_way"`

	// 联系方式值
	ContactValue *string `json:"contact_value,omitempty" xml:"contact_value"`

	// 国家码
	AreaCode *string `json:"area_code,omitempty" xml:"area_code"`

	// 验证序列号
	VerifiedId *string `json:"verified_id,omitempty" xml:"verified_id"`
}

func (o ContactV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContactV2 struct{}"
	}

	return strings.Join([]string{"ContactV2", string(data)}, " ")
}
