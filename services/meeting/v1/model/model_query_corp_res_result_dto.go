package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询企业资源的返回结果
type QueryCorpResResultDto struct {
	CorpVcRes *QueryCorpVcResResultDto `json:"corpVcRes,omitempty" xml:"corpVcRes"`
}

func (o QueryCorpResResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryCorpResResultDto struct{}"
	}

	return strings.Join([]string{"QueryCorpResResultDto", string(data)}, " ")
}
