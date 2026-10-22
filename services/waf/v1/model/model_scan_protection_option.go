package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScanProtectionOption struct {
	HighFrequencyScanBlockingConfig *HighFrequencyScanBlockingConfig `json:"high_frequency_scan_blocking_config,omitempty"`

	DirectoryTraversalProtectionConfig *DirectoryTraversalProtectionConfig `json:"directory_traversal_protection_config,omitempty"`
}

func (o ScanProtectionOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScanProtectionOption struct{}"
	}

	return strings.Join([]string{"ScanProtectionOption", string(data)}, " ")
}
