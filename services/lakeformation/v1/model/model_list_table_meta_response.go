package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTableMetaResponse struct {
	TableMetas *[]TableMeta `json:"table_metas,omitempty"`

	PageInfo       *PagedInfo `json:"page_info,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListTableMetaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableMetaResponse struct{}"
	}

	return strings.Join([]string{"ListTableMetaResponse", string(data)}, " ")
}
