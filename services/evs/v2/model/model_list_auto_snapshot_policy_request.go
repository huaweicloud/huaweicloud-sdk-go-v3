package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutoSnapshotPolicyRequest Request Object
type ListAutoSnapshotPolicyRequest struct {

	// 自动快照策略ID。
	Id *string `json:"id,omitempty"`

	// 自动快照策略名称。
	Name *string `json:"name,omitempty"`

	// 自动快照策略的标签的key/value对，例如[{\"key1\":\"value1\"}]。
	Tags *string `json:"tags,omitempty"`

	// 自动快照策略的标签的key。
	TagKey *string `json:"tag_key,omitempty"`

	// 自动快照策略的企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 返回结果个数限制。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 返回结果按该关键字排序，支持id，created_at等关键字，默认为“created_at”。
	SortKey *string `json:"sort_key,omitempty"`

	// 返回结果按照降序或升序排列，默认为“desc”。 降序：desc 升序：asc
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListAutoSnapshotPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoSnapshotPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListAutoSnapshotPolicyRequest", string(data)}, " ")
}
