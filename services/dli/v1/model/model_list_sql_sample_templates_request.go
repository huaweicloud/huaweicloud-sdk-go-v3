package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlSampleTemplatesRequest Request Object
type ListSqlSampleTemplatesRequest struct {
}

func (o ListSqlSampleTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlSampleTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListSqlSampleTemplatesRequest", string(data)}, " ")
}
