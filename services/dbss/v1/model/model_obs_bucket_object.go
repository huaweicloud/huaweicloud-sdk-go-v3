package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObsBucketObject struct {

	// 桶名称
	BucketName *string `json:"bucket_name,omitempty"`

	// 桶所在的区域
	Location *string `json:"location,omitempty"`
}

func (o ObsBucketObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsBucketObject struct{}"
	}

	return strings.Join([]string{"ObsBucketObject", string(data)}, " ")
}
