package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStoredQueryRequest Request Object
type DeleteStoredQueryRequest struct {

	// 查询ID
	QueryId string `json:"query_id"`
}

func (o DeleteStoredQueryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStoredQueryRequest struct{}"
	}

	return strings.Join([]string{"DeleteStoredQueryRequest", string(data)}, " ")
}
