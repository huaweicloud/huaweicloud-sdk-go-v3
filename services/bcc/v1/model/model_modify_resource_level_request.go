package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyResourceLevelRequest Request Object
type ModifyResourceLevelRequest struct {

	// 账户ID
	DomainId string `json:"domain_id"`

	// 资源等级ID
	LevelId string `json:"level_id"`

	Body *ModifyResourceLevelBody `json:"body,omitempty"`
}

func (o ModifyResourceLevelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyResourceLevelRequest struct{}"
	}

	return strings.Join([]string{"ModifyResourceLevelRequest", string(data)}, " ")
}
