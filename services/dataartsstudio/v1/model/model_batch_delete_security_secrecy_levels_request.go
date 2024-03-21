package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSecuritySecrecyLevelsRequest Request Object
type BatchDeleteSecuritySecrecyLevelsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *BatchDeleteSecrecyLevelDto `json:"body,omitempty"`
}

func (o BatchDeleteSecuritySecrecyLevelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSecuritySecrecyLevelsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteSecuritySecrecyLevelsRequest", string(data)}, " ")
}
