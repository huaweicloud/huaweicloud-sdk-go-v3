package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteV2WorkspaceResponse Response Object
type DeleteV2WorkspaceResponse struct {

	// **参数解释**：工作空间ID。 **取值范围**：长度为1~36的英文字母、数字和中划线的组合。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteV2WorkspaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteV2WorkspaceResponse struct{}"
	}

	return strings.Join([]string{"DeleteV2WorkspaceResponse", string(data)}, " ")
}
