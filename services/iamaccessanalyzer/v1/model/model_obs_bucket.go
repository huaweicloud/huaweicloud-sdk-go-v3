package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObsBucket struct {

	// 此ACL xml文件的string格式
	BucketAcl *string `json:"bucket_acl,omitempty"`

	// 此策略的json格式策略文档。
	BucketPolicy *string `json:"bucket_policy,omitempty"`
}

func (o ObsBucket) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsBucket struct{}"
	}

	return strings.Join([]string{"ObsBucket", string(data)}, " ")
}
