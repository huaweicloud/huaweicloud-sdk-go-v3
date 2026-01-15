package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesCloudStorageOptions 云存储设置项。
type PoliciesCloudStorageOptions struct {

	// 挂盘路径。json格式，长度不能超过500000个字符。
	CloudStorageRule *string `json:"cloud_storage_rule,omitempty"`
}

func (o PoliciesCloudStorageOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesCloudStorageOptions struct{}"
	}

	return strings.Join([]string{"PoliciesCloudStorageOptions", string(data)}, " ")
}
