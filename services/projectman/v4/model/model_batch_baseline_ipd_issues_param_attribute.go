package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchBaselineIpdIssuesParamAttribute 修改的工作项属性。
type BatchBaselineIpdIssuesParamAttribute struct {

	// 工作项基线标识。
	Baseline string `json:"baseline"`
}

func (o BatchBaselineIpdIssuesParamAttribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchBaselineIpdIssuesParamAttribute struct{}"
	}

	return strings.Join([]string{"BatchBaselineIpdIssuesParamAttribute", string(data)}, " ")
}
