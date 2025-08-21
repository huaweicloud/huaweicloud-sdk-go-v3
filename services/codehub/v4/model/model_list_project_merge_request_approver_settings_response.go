package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectMergeRequestApproverSettingsResponse Response Object
type ListProjectMergeRequestApproverSettingsResponse struct {
	Body           *[]MergeRequestApproverSettingResultDto `json:"body,omitempty"`
	HttpStatusCode int                                     `json:"-"`
}

func (o ListProjectMergeRequestApproverSettingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectMergeRequestApproverSettingsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectMergeRequestApproverSettingsResponse", string(data)}, " ")
}
