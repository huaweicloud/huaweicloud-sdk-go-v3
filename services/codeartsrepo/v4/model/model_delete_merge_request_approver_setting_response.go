package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMergeRequestApproverSettingResponse Response Object
type DeleteMergeRequestApproverSettingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMergeRequestApproverSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMergeRequestApproverSettingResponse struct{}"
	}

	return strings.Join([]string{"DeleteMergeRequestApproverSettingResponse", string(data)}, " ")
}
