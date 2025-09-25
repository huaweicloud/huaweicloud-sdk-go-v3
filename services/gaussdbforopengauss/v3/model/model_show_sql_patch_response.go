package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSqlPatchResponse Response Object
type ShowSqlPatchResponse struct {

	// **参数解释**: PATCH名称。 **取值范围**: 不涉及。
	PatchName *string `json:"patch_name,omitempty"`

	// **参数解释**: PATCH内容（Hint文本）。对于abort 类型的SQL PATCH此字段为空。 **取值范围**: 不涉及。
	Hint *string `json:"hint,omitempty"`

	// **参数解释**: PATCH状态。 **取值范围**: - enabled：表示SQL PATCH生效。 - disabled：表示SQL PATCH失效。
	PatchStatus    *string `json:"patch_status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSqlPatchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlPatchResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlPatchResponse", string(data)}, " ")
}
