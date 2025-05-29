package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResultValueListTaskResulDetailtVo struct {

	// 起始记录数 大于 实际总条数时， 值为0， 分页请求才有此值
	Total *int32 `json:"total,omitempty"`

	// 业务失败的提示内容，对内接口才有此值
	Reason *string `json:"reason,omitempty"`

	PageSize *int32 `json:"page_size,omitempty"`

	PageNo *int32 `json:"page_no,omitempty"`

	Value *TaskResultDetailVo `json:"value,omitempty"`
}

func (o ResultValueListTaskResulDetailtVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultValueListTaskResulDetailtVo struct{}"
	}

	return strings.Join([]string{"ResultValueListTaskResulDetailtVo", string(data)}, " ")
}
