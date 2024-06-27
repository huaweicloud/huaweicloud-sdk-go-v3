package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CaseIdAndTypeInfo 用例列表信息
type CaseIdAndTypeInfo struct {

	// 用例id
	Id *string `json:"id,omitempty"`

	// 用例类型, 对应ServiceType
	Type *string `json:"type,omitempty"`
}

func (o CaseIdAndTypeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaseIdAndTypeInfo struct{}"
	}

	return strings.Join([]string{"CaseIdAndTypeInfo", string(data)}, " ")
}
