package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObsBuckets struct {

	// 桶名称。
	BucketName string `json:"bucket_name"`

	// 创建时间,格式为时间戳（毫秒级）。
	CreateTime int64 `json:"create_time"`
}

func (o ObsBuckets) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsBuckets struct{}"
	}

	return strings.Join([]string{"ObsBuckets", string(data)}, " ")
}
