package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelSyncIamUserResponse Response Object
type CancelSyncIamUserResponse struct {

	// 请求下发的结果
	State          *string `json:"state,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelSyncIamUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSyncIamUserResponse struct{}"
	}

	return strings.Join([]string{"CancelSyncIamUserResponse", string(data)}, " ")
}
