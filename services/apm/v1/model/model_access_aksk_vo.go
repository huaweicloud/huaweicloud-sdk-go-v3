package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ak/sk数据模型。
type AccessAkskVo struct {

	// ak/sk的id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// ak/sk的生成时间
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create"`

	// ak/sk的修改时间
	GmtModify *string `json:"gmt_modify,omitempty" xml:"gmt_modify"`

	// 内部租户id
	InnerDomainId *int32 `json:"inner_domain_id,omitempty" xml:"inner_domain_id"`

	// 生成的ak
	Ak *string `json:"ak,omitempty" xml:"ak"`

	// 生成的sk
	Sk *string `json:"sk,omitempty" xml:"sk"`

	// ak/sk的状态
	Status *string `json:"status,omitempty" xml:"status"`

	// ak/sk的描述信息
	Descp *string `json:"descp,omitempty" xml:"descp"`
}

func (o AccessAkskVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessAkskVo struct{}"
	}

	return strings.Join([]string{"AccessAkskVo", string(data)}, " ")
}
