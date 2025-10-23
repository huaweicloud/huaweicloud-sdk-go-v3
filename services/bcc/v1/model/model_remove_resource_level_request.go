package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveResourceLevelRequest Request Object
type RemoveResourceLevelRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	// 资源等级ID
	LevelId string `json:"level_id"`
}

func (o RemoveResourceLevelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveResourceLevelRequest struct{}"
	}

	return strings.Join([]string{"RemoveResourceLevelRequest", string(data)}, " ")
}
