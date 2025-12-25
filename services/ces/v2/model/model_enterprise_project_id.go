package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnterpriseProjectId **参数解释**： 企业项目ID。     **约束限制**： 不涉及。 **取值范围**： 只能包含小写字母、数字、“-”。0 代表默认企业项目ID         **默认取值**： 不涉及。
type EnterpriseProjectId struct {
}

func (o EnterpriseProjectId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseProjectId struct{}"
	}

	return strings.Join([]string{"EnterpriseProjectId", string(data)}, " ")
}
