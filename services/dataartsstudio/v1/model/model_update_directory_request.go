package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDirectoryRequest Request Object
type UpdateDirectoryRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *DirectoryVo `json:"body,omitempty"`
}

func (o UpdateDirectoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDirectoryRequest struct{}"
	}

	return strings.Join([]string{"UpdateDirectoryRequest", string(data)}, " ")
}
