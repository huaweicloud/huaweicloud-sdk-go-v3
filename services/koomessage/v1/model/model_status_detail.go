package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatusDetail struct {

	// 运营商类型。  - cmcc：中国移动 - cucc：中国联通 - ctcc：中国电信 - oversea：港澳台及国外 - unknown：未知
	Carrier *string `json:"carrier,omitempty"`

	// 模板状态： - 0：正常可用  - 1：审核中  - 2：审核不通过  - 3：模板已禁用  - 4：模板不存在  - 5：模板已过期
	Status *int32 `json:"status,omitempty"`

	// 对模板状态的描述。  > 若状态是审核不通过或被禁用，描述表示的是不通过或禁用的原因。
	Desc *string `json:"desc,omitempty"`
}

func (o StatusDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusDetail struct{}"
	}

	return strings.Join([]string{"StatusDetail", string(data)}, " ")
}
