package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProjectMergeRequestApproverSettingResponse Response Object
type DeleteProjectMergeRequestApproverSettingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteProjectMergeRequestApproverSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProjectMergeRequestApproverSettingResponse struct{}"
	}

	return strings.Join([]string{"DeleteProjectMergeRequestApproverSettingResponse", string(data)}, " ")
}
