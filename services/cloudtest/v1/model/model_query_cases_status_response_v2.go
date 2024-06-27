package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryCasesStatusResponseV2 用例状态数据
type QueryCasesStatusResponseV2 struct {
	CasesStatusJA *[]interface{} `json:"casesStatusJA,omitempty"`
}

func (o QueryCasesStatusResponseV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCasesStatusResponseV2 struct{}"
	}

	return strings.Join([]string{"QueryCasesStatusResponseV2", string(data)}, " ")
}
