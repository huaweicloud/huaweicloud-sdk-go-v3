package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询企业资源的返回结果
type QueryCorpResResultDto struct {
	CorpVcRes *QueryCorpVcResResultDto `json:"corpVcRes,omitempty"`
}

func (o QueryCorpResResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCorpResResultDto struct{}"
	}

	return strings.Join([]string{"QueryCorpResResultDto", string(data)}, " ")
}
