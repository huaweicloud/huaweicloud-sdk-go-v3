package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlJobSystemDefendRulesRequest Request Object
type ListSqlJobSystemDefendRulesRequest struct {
}

func (o ListSqlJobSystemDefendRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlJobSystemDefendRulesRequest struct{}"
	}

	return strings.Join([]string{"ListSqlJobSystemDefendRulesRequest", string(data)}, " ")
}
