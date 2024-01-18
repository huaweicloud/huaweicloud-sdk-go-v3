package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DriftDetail 漂移详细信息。
type DriftDetail struct {

	// 管理纳管账号ID。
	ManageAccountId *string `json:"manage_account_id,omitempty"`

	// 漂移类型。
	DriftType *string `json:"drift_type,omitempty"`

	// 漂移发生的纳管账号ID或注册OU ID。
	DriftTargetId *string `json:"drift_target_id,omitempty"`

	// 漂移目标类型，类型有accountId和policyId。
	DriftTargetType *string `json:"drift_target_type,omitempty"`

	// 漂移信息。
	DriftMessage *string `json:"drift_message,omitempty"`

	// 父注册OU ID。
	ParentOrganizationUnitId *string `json:"parent_organization_unit_id,omitempty"`
}

func (o DriftDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DriftDetail struct{}"
	}

	return strings.Join([]string{"DriftDetail", string(data)}, " ")
}
