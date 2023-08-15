package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListDcDsResponse Response Object
type BatchListDcDsResponse struct {

	// 总记录数
	Count *int64 `json:"count,omitempty"`

	PageInfo *PageInfoDto `json:"page_info,omitempty"`

	// 每页记录数
	Datasources    *[]QueryDcDsBriefRespDto `json:"datasources,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o BatchListDcDsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListDcDsResponse struct{}"
	}

	return strings.Join([]string{"BatchListDcDsResponse", string(data)}, " ")
}
