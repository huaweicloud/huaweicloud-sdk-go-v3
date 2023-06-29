package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPassphraseResponse Response Object
type ShowPassphraseResponse struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 安全传输通道证书passphrase
	Passphrase     *string `json:"passphrase,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPassphraseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPassphraseResponse struct{}"
	}

	return strings.Join([]string{"ShowPassphraseResponse", string(data)}, " ")
}
