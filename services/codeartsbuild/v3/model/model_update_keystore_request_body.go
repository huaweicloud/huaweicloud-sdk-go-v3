package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKeystoreRequestBody job_ids列表
type UpdateKeystoreRequestBody struct {

	// 文件ID
	Id string `json:"id"`

	// 文件描述
	Description string `json:"description"`

	// 文件名
	KeystoreName string `json:"keystore_name"`

	// 是否租户内共享，1-共享，0-不共享
	Share int32 `json:"share"`
}

func (o UpdateKeystoreRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeystoreRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateKeystoreRequestBody", string(data)}, " ")
}
