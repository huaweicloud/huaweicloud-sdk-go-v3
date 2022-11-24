package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Pager struct {

	// 页码
	Page *int32 `json:"page,omitempty"`

	// 每页大小
	PageSize *int32 `json:"page_size,omitempty"`

	// 总条数
	TotalRows *int32 `json:"total_rows,omitempty"`
}

func (o Pager) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Pager struct{}"
	}

	return strings.Join([]string{"Pager", string(data)}, " ")
}
