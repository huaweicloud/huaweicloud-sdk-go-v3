package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessAkskVo ak/sk数据模型。
type AccessAkskVo struct {

	// ak/sk的id。
	Id *int32 `json:"id,omitempty"`

	// ak/sk的生成时间。
	GmtCreate *string `json:"gmt_create,omitempty"`

	// ak/sk的修改时间。
	GmtModify *string `json:"gmt_modify,omitempty"`

	// 内部租户id。
	InnerDomainId *int32 `json:"inner_domain_id,omitempty"`

	// 生成的ak。
	Ak *string `json:"ak,omitempty"`

	// 生成的sk。
	Sk *string `json:"sk,omitempty"`

	// ak/sk的状态。
	Status *string `json:"status,omitempty"`

	// ak/sk的描述信息。
	Descp *string `json:"descp,omitempty"`

	// ak/sk的生成时间戳。
	GmtCreateTimestamp *int64 `json:"gmt_create_timestamp,omitempty"`

	// ak/sk的修改时间戳。
	GmtModifyTimestamp *int64 `json:"gmt_modify_timestamp,omitempty"`
}

func (o AccessAkskVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessAkskVo struct{}"
	}

	return strings.Join([]string{"AccessAkskVo", string(data)}, " ")
}
