package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartLogAutoBackupPolicyResponse Response Object
type StartLogAutoBackupPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartLogAutoBackupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartLogAutoBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"StartLogAutoBackupPolicyResponse", string(data)}, " ")
}
