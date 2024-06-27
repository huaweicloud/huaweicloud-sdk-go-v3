package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CaseAwInstance struct {

	// 测试用例前置步骤
	SetupAwList *[]AwInstance `json:"setup_aw_list,omitempty"`

	// 测试步骤
	TeardownAwList *[]AwInstance `json:"teardown_aw_list,omitempty"`

	// 测试用例后置不走
	TestAwList *[]AwInstance `json:"test_aw_list,omitempty"`
}

func (o CaseAwInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaseAwInstance struct{}"
	}

	return strings.Join([]string{"CaseAwInstance", string(data)}, " ")
}
