package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KmsResourceInfo struct {

	// KMS密钥数量
	TotalNum int32 `json:"total_num"`

	Result *KmsInfo `json:"result"`
}

func (o KmsResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KmsResourceInfo struct{}"
	}

	return strings.Join([]string{"KmsResourceInfo", string(data)}, " ")
}
