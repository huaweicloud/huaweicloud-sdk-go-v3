package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDocumentAtomicsResponse Response Object
type ListDocumentAtomicsResponse struct {

	// 原子能力列表
	Data *[]AtomicModel `json:"data,omitempty"`

	// 总条数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDocumentAtomicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDocumentAtomicsResponse struct{}"
	}

	return strings.Join([]string{"ListDocumentAtomicsResponse", string(data)}, " ")
}
