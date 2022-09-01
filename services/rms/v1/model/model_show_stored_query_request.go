package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowStoredQueryRequest struct {

	// 查询ID
	QueryId string `json:"query_id" xml:"query_id"`
}

func (o ShowStoredQueryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStoredQueryRequest struct{}"
	}

	return strings.Join([]string{"ShowStoredQueryRequest", string(data)}, " ")
}
