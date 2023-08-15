package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlSampleTemplatesRequest Request Object
type ShowSqlSampleTemplatesRequest struct {
}

func (o ShowSqlSampleTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlSampleTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlSampleTemplatesRequest", string(data)}, " ")
}
