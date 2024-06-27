package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChEnlargeVolumeRequestBody 磁盘扩容请求体。
type ChEnlargeVolumeRequestBody struct {

	// 磁盘容量。取值范围：50GB~32000GB。
	SizeInGb int32 `json:"size_in_gb"`
}

func (o ChEnlargeVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChEnlargeVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"ChEnlargeVolumeRequestBody", string(data)}, " ")
}
