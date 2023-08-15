package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlaybookVersionRequest Request Object
type ShowPlaybookVersionRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// version Id value
	VersionId string `json:"version_id"`
}

func (o ShowPlaybookVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowPlaybookVersionRequest", string(data)}, " ")
}
