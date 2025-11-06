package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartRemoteMirrorSynchronizationResponse Response Object
type StartRemoteMirrorSynchronizationResponse struct {

	// **参数解释：** 同步任务Id。
	Jid            *string `json:"jid,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartRemoteMirrorSynchronizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartRemoteMirrorSynchronizationResponse struct{}"
	}

	return strings.Join([]string{"StartRemoteMirrorSynchronizationResponse", string(data)}, " ")
}
