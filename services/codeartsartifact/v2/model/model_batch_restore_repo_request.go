package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRestoreRepoRequest Request Object
type BatchRestoreRepoRequest struct {
	Body *[]IdeTrashArtifactModel `json:"body,omitempty"`
}

func (o BatchRestoreRepoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestoreRepoRequest struct{}"
	}

	return strings.Join([]string{"BatchRestoreRepoRequest", string(data)}, " ")
}
