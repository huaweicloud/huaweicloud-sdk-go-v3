package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFilePushPermissionResponse Response Object
type CreateFilePushPermissionResponse struct {

	// **参数解释：** 仓库文件推送权限ID。 **取值范围：** 1-2147483647
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 目录路径或通配符。 **取值范围：** 不涉及。
	Path *string `json:"path,omitempty"`

	// **参数解释：** 事件列表。 **取值范围：** 不涉及。
	Actions        *[]RepositoryProtectedActionDto `json:"actions,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o CreateFilePushPermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFilePushPermissionResponse struct{}"
	}

	return strings.Join([]string{"CreateFilePushPermissionResponse", string(data)}, " ")
}
