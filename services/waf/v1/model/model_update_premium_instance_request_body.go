package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePremiumInstanceRequestBody 独享引擎操作
type UpdatePremiumInstanceRequestBody struct {

	// **参数解释：** 独享引擎操作名称，目前支持 upgrade（升级） ，rollback（升级后回滚），security_groups（切换安全组） **取值范围：** - upgrade - rollback - security_groups
	Action string `json:"action"`

	// **参数解释：**  具体的请求参数，操作为upgrade（升级） ，rollback（升级后回滚）时无需填写，操作为 security_groups（切换安全组）时，参数为安全组的ip数组 **取值范围：** 不涉及
	Params *[]string `json:"params,omitempty"`
}

func (o UpdatePremiumInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePremiumInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePremiumInstanceRequestBody", string(data)}, " ")
}
