package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDirectoryRequest Request Object
type CreateDirectoryRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *DirectoryVo `json:"body,omitempty"`
}

func (o CreateDirectoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDirectoryRequest struct{}"
	}

	return strings.Join([]string{"CreateDirectoryRequest", string(data)}, " ")
}
