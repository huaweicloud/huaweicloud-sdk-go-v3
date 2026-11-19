package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlPatchResponse Response Object
type ListSqlPatchResponse struct {

	// **参数解释**: SQL Patch信息列表。
	SqlPatchInfos *[]SqlPatchInfosResult `json:"sql_patch_infos,omitempty"`

	// **参数解释**: 返回的SQL Patch数量。 **取值范围**: 不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSqlPatchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlPatchResponse struct{}"
	}

	return strings.Join([]string{"ListSqlPatchResponse", string(data)}, " ")
}
