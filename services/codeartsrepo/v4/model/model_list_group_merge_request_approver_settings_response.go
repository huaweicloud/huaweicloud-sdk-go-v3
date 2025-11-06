package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupMergeRequestApproverSettingsResponse Response Object
type ListGroupMergeRequestApproverSettingsResponse struct {
	Body           *[]MergeRequestApproverSettingResultDto `json:"body,omitempty"`
	HttpStatusCode int                                     `json:"-"`
}

func (o ListGroupMergeRequestApproverSettingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupMergeRequestApproverSettingsResponse struct{}"
	}

	return strings.Join([]string{"ListGroupMergeRequestApproverSettingsResponse", string(data)}, " ")
}
