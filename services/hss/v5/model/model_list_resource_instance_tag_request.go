package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceInstanceTagRequest Request Object
type ListResourceInstanceTagRequest struct {

	// **参数解释**: 由标签管理服务定义的资源类别，企业主机安全服务调用此接口时资源类别为hss **约束限制**: 不涉及 **取值范围**: - hss：hss **默认取值**: 不涉及
	ResourceType string `json:"resource_type"`

	// **参数解释**: 资源ID **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	ResourceId string `json:"resource_id"`
}

func (o ListResourceInstanceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceInstanceTagRequest struct{}"
	}

	return strings.Join([]string{"ListResourceInstanceTagRequest", string(data)}, " ")
}
