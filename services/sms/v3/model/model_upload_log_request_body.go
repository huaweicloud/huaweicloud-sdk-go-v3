package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadLogRequestBody 上传日志请求体
type UploadLogRequestBody struct {

	// 指定桶名称
	LogBucket string `json:"log_bucket"`

	// 指定有效期
	LogExpire int32 `json:"log_expire"`
}

func (o UploadLogRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadLogRequestBody struct{}"
	}

	return strings.Join([]string{"UploadLogRequestBody", string(data)}, " ")
}
