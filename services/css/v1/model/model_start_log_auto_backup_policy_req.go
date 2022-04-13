package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartLogAutoBackupPolicyReq struct {
	// 备份开始时间。

	Period string `json:"period"`
}

func (o StartLogAutoBackupPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartLogAutoBackupPolicyReq struct{}"
	}

	return strings.Join([]string{"StartLogAutoBackupPolicyReq", string(data)}, " ")
}
