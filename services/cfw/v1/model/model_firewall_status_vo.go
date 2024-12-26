package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FirewallStatusVo struct {

	// 可防护eip数量
	AvailableEipCount *int32 `json:"available_eip_count,omitempty"`

	// 是否超出eip数量限制
	BeyondMaxCount *bool `json:"beyond_max_count,omitempty"`

	// 已防护eip数量
	EipProtectedSelf *int32 `json:"eip_protected_self,omitempty"`

	// eip总数
	EipTotal *int32 `json:"eip_total,omitempty"`

	// 未防护eip数量
	EipUnProtected *int32 `json:"eip_un_protected,omitempty"`

	// 防护对象id
	ObjectId *string `json:"object_id,omitempty"`

	// 是否开启新增eip自动防护，1；是，0：否
	Status *int32 `json:"status,omitempty"`
}

func (o FirewallStatusVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FirewallStatusVo struct{}"
	}

	return strings.Join([]string{"FirewallStatusVo", string(data)}, " ")
}
