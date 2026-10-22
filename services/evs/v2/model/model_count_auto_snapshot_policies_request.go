package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountAutoSnapshotPoliciesRequest Request Object
type CountAutoSnapshotPoliciesRequest struct {

	// 自动快照策略ID。
	Id *string `json:"id,omitempty"`

	// 自动快照策略名称。
	Name *string `json:"name,omitempty"`

	// 自动快照策略的标签的key。
	TagKey *string `json:"tag_key,omitempty"`

	// 自动快照策略的标签的key/value对，例如[{\"key1\":\"value1\"}]。
	Tags *string `json:"tags,omitempty"`

	// 自动快照策略的企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CountAutoSnapshotPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountAutoSnapshotPoliciesRequest struct{}"
	}

	return strings.Join([]string{"CountAutoSnapshotPoliciesRequest", string(data)}, " ")
}
