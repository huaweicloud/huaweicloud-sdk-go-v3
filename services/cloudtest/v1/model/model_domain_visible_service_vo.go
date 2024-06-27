package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DomainVisibleServiceVo 实际的数据类型：单个对象，集合 或 NULL
type DomainVisibleServiceVo struct {

	// 第三方服务名
	Name *string `json:"name,omitempty"`

	// 第三方服务类型
	Type *int32 `json:"type,omitempty"`

	// 第三方服务执行方式（0：普通TestHub，1：对接八爪鱼TestHub）
	ExecuteType *int32 `json:"execute_type,omitempty"`
}

func (o DomainVisibleServiceVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainVisibleServiceVo struct{}"
	}

	return strings.Join([]string{"DomainVisibleServiceVo", string(data)}, " ")
}
