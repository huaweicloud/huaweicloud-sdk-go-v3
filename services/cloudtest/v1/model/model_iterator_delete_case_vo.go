package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IteratorDeleteCaseVo 实际的数据类型：单个对象，集合 或 NULL
type IteratorDeleteCaseVo struct {

	// 操作的id, 由projectUuid + - + iteratorUri + - + caseId 组成
	Id *string `json:"id,omitempty"`

	// 操作名称
	Name *string `json:"name,omitempty"`
}

func (o IteratorDeleteCaseVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IteratorDeleteCaseVo struct{}"
	}

	return strings.Join([]string{"IteratorDeleteCaseVo", string(data)}, " ")
}
