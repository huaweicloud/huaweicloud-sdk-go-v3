package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultValueListProjectTestCaseFieldVo 请求的返回的数据对象
type ResultValueListProjectTestCaseFieldVo struct {
	Value *[]ProjectTestCaseFieldVo `json:"value,omitempty"`
}

func (o ResultValueListProjectTestCaseFieldVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultValueListProjectTestCaseFieldVo struct{}"
	}

	return strings.Join([]string{"ResultValueListProjectTestCaseFieldVo", string(data)}, " ")
}
