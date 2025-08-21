package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLastPushEventInRepositoryResponse Response Object
type ShowLastPushEventInRepositoryResponse struct {

	// **参数解释：** 分支或者tag名称。 **取值范围：** 不涉及。
	Ref *string `json:"ref,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	Repository     *RepositoryBriefDto `json:"repository,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowLastPushEventInRepositoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLastPushEventInRepositoryResponse struct{}"
	}

	return strings.Join([]string{"ShowLastPushEventInRepositoryResponse", string(data)}, " ")
}
