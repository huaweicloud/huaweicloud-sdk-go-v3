package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlDetailResponse Response Object
type ShowSqlDetailResponse struct {
	Sql            *SqlDetailBean `json:"sql,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowSqlDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlDetailResponse", string(data)}, " ")
}
