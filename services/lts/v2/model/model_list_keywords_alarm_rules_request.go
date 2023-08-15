package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKeywordsAlarmRulesRequest Request Object
type ListKeywordsAlarmRulesRequest struct {
}

func (o ListKeywordsAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeywordsAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"ListKeywordsAlarmRulesRequest", string(data)}, " ")
}
