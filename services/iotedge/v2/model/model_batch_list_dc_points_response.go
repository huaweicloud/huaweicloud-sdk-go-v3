package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchListDcPointsResponse struct {

	// 总记录数
	Count *int64 `json:"count,omitempty"`

	PageInfo *PageInfoDto `json:"page_info,omitempty"`

	// 每页记录数
	Points         *[]CreateDcPointRespDto `json:"points,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o BatchListDcPointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListDcPointsResponse struct{}"
	}

	return strings.Join([]string{"BatchListDcPointsResponse", string(data)}, " ")
}
