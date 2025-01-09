package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBucketCredentialReq 生成资源访问凭据。
type CreateBucketCredentialReq struct {

	// 文件完整名称,不允许包含如下字符:^;|~`{}[]<>。
	FileName string `json:"file_name"`
}

func (o CreateBucketCredentialReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBucketCredentialReq struct{}"
	}

	return strings.Join([]string{"CreateBucketCredentialReq", string(data)}, " ")
}
