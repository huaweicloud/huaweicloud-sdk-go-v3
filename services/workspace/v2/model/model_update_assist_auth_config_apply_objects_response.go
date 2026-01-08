package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssistAuthConfigApplyObjectsResponse Response Object
type UpdateAssistAuthConfigApplyObjectsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAssistAuthConfigApplyObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssistAuthConfigApplyObjectsResponse struct{}"
	}

	return strings.Join([]string{"UpdateAssistAuthConfigApplyObjectsResponse", string(data)}, " ")
}
