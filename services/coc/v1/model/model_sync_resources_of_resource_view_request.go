package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncResourcesOfResourceViewRequest Request Object
type SyncResourcesOfResourceViewRequest struct {

	// **参数解释**: 视图ID。 **约束限制**: 不涉及。 **取值范围**: 自定义生成，长度1~32之间。 **默认取值**: 不涉及。
	Id string `json:"id"`
}

func (o SyncResourcesOfResourceViewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncResourcesOfResourceViewRequest struct{}"
	}

	return strings.Join([]string{"SyncResourcesOfResourceViewRequest", string(data)}, " ")
}
