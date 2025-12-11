package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectTagsRequest Request Object
type ListProjectTagsRequest struct {

	// **参数解释**: 由标签管理服务定义的资源类别，企业主机安全服务调用此接口时资源类别为hss。 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位。 **默认取值**: hss
	ResourceType string `json:"resource_type"`
}

func (o ListProjectTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectTagsRequest", string(data)}, " ")
}
