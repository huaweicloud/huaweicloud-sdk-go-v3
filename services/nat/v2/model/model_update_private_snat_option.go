package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePrivateSnatOption 更新SNAT规则的请求体。
type UpdatePrivateSnatOption struct {

	// 中转IP的ID的列表。
	TransitIpIds *[]string `json:"transit_ip_ids,omitempty"`

	// SNAT规则的描述。长度范围小于等于255个字符，不能包含“<”和“>”。
	Description *string `json:"description,omitempty"`
}

func (o UpdatePrivateSnatOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateSnatOption struct{}"
	}

	return strings.Join([]string{"UpdatePrivateSnatOption", string(data)}, " ")
}
