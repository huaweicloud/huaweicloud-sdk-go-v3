package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultValueListDomainVisibleServiceVo 请求的返回的数据对象
type ResultValueListDomainVisibleServiceVo struct {

	// 实际的数据类型：单个对象，集合 或 NULL
	Value *[]DomainVisibleServiceVo `json:"value,omitempty"`
}

func (o ResultValueListDomainVisibleServiceVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultValueListDomainVisibleServiceVo struct{}"
	}

	return strings.Join([]string{"ResultValueListDomainVisibleServiceVo", string(data)}, " ")
}
