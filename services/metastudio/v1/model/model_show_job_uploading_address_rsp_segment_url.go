package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobUploadingAddressRspSegmentUrl 分句上传任务的上传地址。
type ShowJobUploadingAddressRspSegmentUrl struct {

	// 音频上传的地址。 > * 通过该obs地址上传时需要设置content-type为audio/wav
	AudioUploadingUrl *[]string `json:"audio_uploading_url,omitempty"`

	// 文本上传的地址。 > * 通过该obs地址上传时需要设置content-type为text/plain
	TxtUploadingUrl *[]string `json:"txt_uploading_url,omitempty"`
}

func (o ShowJobUploadingAddressRspSegmentUrl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobUploadingAddressRspSegmentUrl struct{}"
	}

	return strings.Join([]string{"ShowJobUploadingAddressRspSegmentUrl", string(data)}, " ")
}
