package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeadLockRelationshipResponse Response Object
type ShowDeadLockRelationshipResponse struct {

	// 会话列表
	ProcessList *[]DeadLockProcess `json:"process_list,omitempty"`

	// 资源列表
	ResourceList   *[]DeadLockResource `json:"resource_list,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowDeadLockRelationshipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockRelationshipResponse struct{}"
	}

	return strings.Join([]string{"ShowDeadLockRelationshipResponse", string(data)}, " ")
}
