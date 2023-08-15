package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimSendReportsMode 报表数据体。
type ListAimSendReportsMode struct {

	// 数据列表。
	DataList *[]AimSendReport `json:"data_list,omitempty"`

	PageInfo *Page `json:"page_info,omitempty"`
}

func (o ListAimSendReportsMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimSendReportsMode struct{}"
	}

	return strings.Join([]string{"ListAimSendReportsMode", string(data)}, " ")
}
