package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpBlacklistVo 一条IP黑名单的导入信息，包括文件名称、生效范围、导入状态、导入时间和导入失败的错误信息。
type IpBlacklistVo struct {

	// IP黑名单的文件名，对应导出时的文件名
	Name *string `json:"name,omitempty"`

	// IP黑名单的生效范围，1表示EIP，2表示NAT
	EffectScope *[]int32 `json:"effect_scope,omitempty"`

	// IP黑名单的导入状态，有三种状态：2：生成中、1：成功、0：失败。
	ImportStatus *int32 `json:"import_status,omitempty"`

	// IP黑名单的导入时间
	ImportTime *int64 `json:"import_time,omitempty"`

	// 导入失败的错误信息
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o IpBlacklistVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpBlacklistVo struct{}"
	}

	return strings.Join([]string{"IpBlacklistVo", string(data)}, " ")
}
