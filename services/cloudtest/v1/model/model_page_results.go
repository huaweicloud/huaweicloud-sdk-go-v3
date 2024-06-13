package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageResults struct {
	PageList *[]VariableRes `json:"page_list,omitempty"`

	PageNo *int32 `json:"page_no,omitempty"`

	PageSize *int32 `json:"page_size,omitempty"`

	TotalPage *int32 `json:"total_page,omitempty"`

	TotalSize *int64 `json:"total_size,omitempty"`
}

func (o PageResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageResults struct{}"
	}

	return strings.Join([]string{"PageResults", string(data)}, " ")
}
