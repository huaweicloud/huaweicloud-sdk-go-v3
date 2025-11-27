package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGitOpsStatusResponse Response Object
type ShowGitOpsStatusResponse struct {
	Metadata *ListMeta `json:"metadata,omitempty"`

	// Pod列表
	Items          *[]Pod `json:"items,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowGitOpsStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGitOpsStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowGitOpsStatusResponse", string(data)}, " ")
}
