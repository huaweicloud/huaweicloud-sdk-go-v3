package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSyncIamUserResponse Response Object
type UpdateSyncIamUserResponse struct {

	// 请求下发的结果
	State          *string `json:"state,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSyncIamUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSyncIamUserResponse struct{}"
	}

	return strings.Join([]string{"UpdateSyncIamUserResponse", string(data)}, " ")
}
