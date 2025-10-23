package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateRepositoryUserGroupResponse Response Object
type AssociateRepositoryUserGroupResponse struct {

	// **参数解释：** 关联结果。 **取值范围：** - success，关联成功。 - error,关联失败。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssociateRepositoryUserGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateRepositoryUserGroupResponse struct{}"
	}

	return strings.Join([]string{"AssociateRepositoryUserGroupResponse", string(data)}, " ")
}
