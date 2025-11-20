package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAgencyRequest Request Object
type DeleteAgencyRequest struct {

	// **参数解释：** purged取值为true时，会同步删除在IAM处创建的premium_waf_svc_trust委托，purged取值为false时，不会同步删除在IAM处创建的premium_waf_svc_trust委托 **约束限制：** 不涉及 **取值范围：** - true - false  **默认取值：** 不涉及
	Purged *bool `json:"purged,omitempty"`

	// **参数解释：** 待删除的代理id **约束限制：** 不涉及 **取值范围：** 从 “查询独享引擎代理”接口的返回结果中，选取需要删除代理的id值 **默认取值：** 不涉及
	RoleIdList *[]string `json:"role_id_list,omitempty"`
}

func (o DeleteAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAgencyRequest struct{}"
	}

	return strings.Join([]string{"DeleteAgencyRequest", string(data)}, " ")
}
