package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestApproverSettingsResponse Response Object
type ListMergeRequestApproverSettingsResponse struct {
	Body           *[]MergeRequestApproverSettingResultDto `json:"body,omitempty"`
	HttpStatusCode int                                     `json:"-"`
}

func (o ListMergeRequestApproverSettingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestApproverSettingsResponse struct{}"
	}

	return strings.Join([]string{"ListMergeRequestApproverSettingsResponse", string(data)}, " ")
}
