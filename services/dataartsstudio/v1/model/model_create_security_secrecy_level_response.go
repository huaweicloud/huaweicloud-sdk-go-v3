package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecuritySecrecyLevelResponse Response Object
type CreateSecuritySecrecyLevelResponse struct {

	// 密级id
	SecrecyLevelId *string `json:"secrecy_level_id,omitempty"`

	// 密级名称
	SecrecyLevelName *string `json:"secrecy_level_name,omitempty"`

	// 密级等级
	SecrecyLevelNumber *int32 `json:"secrecy_level_number,omitempty"`

	// 密级描述
	Description *string `json:"description,omitempty"`

	// 创建者
	CreatedBy *string `json:"created_by,omitempty"`

	// 创建时间
	CreatedAt *int64 `json:"created_at,omitempty"`

	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty"`

	// 更新时间
	UpdatedAt *int64 `json:"updated_at,omitempty"`

	// DataArts实例ID
	InstanceId     *string `json:"instance_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSecuritySecrecyLevelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecuritySecrecyLevelResponse struct{}"
	}

	return strings.Join([]string{"CreateSecuritySecrecyLevelResponse", string(data)}, " ")
}
