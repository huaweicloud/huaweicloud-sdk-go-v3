package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowPlaybookStatisticsRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`
}

func (o ShowPlaybookStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowPlaybookStatisticsRequest", string(data)}, " ")
}
