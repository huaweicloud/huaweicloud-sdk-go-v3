package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpBlacklistImportDto 导入IP黑名单的消息body体，包括了导入的方式、生效范围和IP黑名单信息。
type IpBlacklistImportDto struct {

	// IP黑名单导入的方式，0表示增量导入，在原来的基础上追加；1表示全量导入，新导入的IP黑名单会覆盖已有的IP黑名单
	AddType *int32 `json:"add_type,omitempty"`

	// IP列表，当前支持不同的IP地址之间使用“,” 、“ ”、“;”、“\\r\\n”、“\\n”、“\\t”等分割符进行分割。
	IpBlacklist *string `json:"ip_blacklist,omitempty"`

	// 生效范围，1表示生效范围为eip, 2表示生效范围为nat, [1 2]表示 生效范围为eip和nat
	EffectScope *[]int32 `json:"effect_scope,omitempty"`
}

func (o IpBlacklistImportDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpBlacklistImportDto struct{}"
	}

	return strings.Join([]string{"IpBlacklistImportDto", string(data)}, " ")
}
