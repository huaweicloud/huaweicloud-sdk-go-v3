package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCsrResponse Response Object
type ListCsrResponse struct {

	// CSR列表，详情请参见CSRResponseBody字段数据结构说明。
	CsrList *[]CsrResponseBody `json:"csr_list,omitempty"`

	// CSR数量。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListCsrResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCsrResponse struct{}"
	}

	return strings.Join([]string{"ListCsrResponse", string(data)}, " ")
}
