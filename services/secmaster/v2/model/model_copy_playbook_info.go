package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Copy playbook
type CopyPlaybookInfo struct {

	// 名称
	Name *string `json:"name,omitempty"`
}

func (o CopyPlaybookInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyPlaybookInfo struct{}"
	}

	return strings.Join([]string{"CopyPlaybookInfo", string(data)}, " ")
}
