package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloudStorageOptions 云存储配置。
type CloudStorageOptions struct {

	// 配置项内容。
	CloudStorageRule *string `json:"cloud_storage_rule,omitempty"`
}

func (o CloudStorageOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudStorageOptions struct{}"
	}

	return strings.Join([]string{"CloudStorageOptions", string(data)}, " ")
}
