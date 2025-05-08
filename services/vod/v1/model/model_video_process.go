package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VideoProcess struct {

	// 需要单独设置时长的HLS起始分片数量。 取值范围：[0,10]，默认值为：0 与hls_init_interval配合使用，设置前面hls_init_count个HLS分片时长。为0表示不单独配置时长。
	HlsInitCount *int32 `json:"hls_init_count,omitempty"`

	// 表示前面hls_init_count个HLS分片的时长。 取值范围：[2,10] ，默认值为：5 hls_init_count不为0时，该字段才起作用。
	HlsInitInterval *int32 `json:"hls_init_interval,omitempty"`
}

func (o VideoProcess) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoProcess struct{}"
	}

	return strings.Join([]string{"VideoProcess", string(data)}, " ")
}
