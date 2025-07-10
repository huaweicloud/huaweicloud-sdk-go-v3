package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferBackupResponse Response Object
type TransferBackupResponse struct {

	// 转储任务结果
	Results        *[]TransferBackupResponseBodyResults `json:"results,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o TransferBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferBackupResponse struct{}"
	}

	return strings.Join([]string{"TransferBackupResponse", string(data)}, " ")
}
