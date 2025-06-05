package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListKeystoreResult struct {

	// 文件名
	KeystoreName *string `json:"keystore_name,omitempty"`

	// 文件ID
	Id *string `json:"id,omitempty"`
}

func (o ListKeystoreResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeystoreResult struct{}"
	}

	return strings.Join([]string{"ListKeystoreResult", string(data)}, " ")
}
