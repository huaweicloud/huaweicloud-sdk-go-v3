package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlertDetail struct {
	DataObject *Alert `json:"data_object,omitempty"`

	// Create time
	CreateTime *string `json:"create_time,omitempty"`

	// Update time
	UpdateTime *string `json:"update_time,omitempty"`

	// Id value
	ProjectId *string `json:"project_id,omitempty"`

	// Id value
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// The name, display only
	Id *string `json:"id,omitempty"`

	// The name, display only
	Type *string `json:"type,omitempty"`

	// The name, display only
	Version *int32 `json:"version,omitempty"`

	// The name, display only
	FormatVersion *int32 `json:"format_version,omitempty"`

	DataclassRef *ShowAlertDetailDataclassRef `json:"dataclass_ref,omitempty"`
}

func (o AlertDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertDetail struct{}"
	}

	return strings.Join([]string{"AlertDetail", string(data)}, " ")
}
