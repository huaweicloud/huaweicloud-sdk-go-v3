package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnterpriseProjectIdResp **参数解释**： 企业项目ID。     **取值范围**： 只能包含小写字母、数字、“-”。
type EnterpriseProjectIdResp struct {
}

func (o EnterpriseProjectIdResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseProjectIdResp struct{}"
	}

	return strings.Join([]string{"EnterpriseProjectIdResp", string(data)}, " ")
}
