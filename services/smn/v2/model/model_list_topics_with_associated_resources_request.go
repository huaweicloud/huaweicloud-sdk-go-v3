package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopicsWithAssociatedResourcesRequest Request Object
type ListTopicsWithAssociatedResourcesRequest struct {

	// 偏移量。  偏移量为一个大于0小于资源总个数的整数，表示查询该偏移量后面的所有的资源，默认值为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的数量限制。  取值范围：1~100，取值一般为10，20，50。功能说明：每页返回的资源个数。默认值为100。
	Limit *int32 `json:"limit,omitempty"`

	// 检索的主题ID，完全匹配。
	TopicId *string `json:"topic_id,omitempty"`

	// 企业项目id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 检索的主题名称，完全匹配。
	Name *string `json:"name,omitempty"`

	// 检索的主题名称，模糊匹配。
	FuzzyName *string `json:"fuzzy_name,omitempty"`

	// 检索的主题显示名。模糊匹配。参数字节长度不能大于192字节。
	FuzzyDisplayName *string `json:"fuzzy_display_name,omitempty"`
}

func (o ListTopicsWithAssociatedResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicsWithAssociatedResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListTopicsWithAssociatedResourcesRequest", string(data)}, " ")
}
