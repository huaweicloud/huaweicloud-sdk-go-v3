package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Information of playbook
type ModifyPlaybookInfo struct {

	// The name, display only
	Name *string `json:"name,omitempty"`

	// The description, display only
	Description *string `json:"description,omitempty"`

	// If is enabled, false for disenabled, true for enabled
	Enabled *bool `json:"enabled,omitempty"`

	// 启用的剧本版本ID
	ActiveVersionId *string `json:"active_version_id,omitempty"`
}

func (o ModifyPlaybookInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPlaybookInfo struct{}"
	}

	return strings.Join([]string{"ModifyPlaybookInfo", string(data)}, " ")
}
