package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DocumentCreateRequestData 文档数据输入
type DocumentCreateRequestData struct {

	// 文档url。目前支持：公网HTTP/HTTPS URL。
	Url string `json:"url"`

	// 文档格式。可选值： docx pdf doc xls xlsx ppt pptx pps ppsx xltx xltm xlsb xlsm txt csv epub webpage 若format与文档实际格式不一致，则返回报错参数错误
	Format string `json:"format"`

	// 当需要审核网页视频时，视频截帧频率间隔，单位为秒，取值范围为1~60s；若不传递默认5s截帧一次
	FrameInterval *int32 `json:"frame_interval,omitempty"`
}

func (o DocumentCreateRequestData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocumentCreateRequestData struct{}"
	}

	return strings.Join([]string{"DocumentCreateRequestData", string(data)}, " ")
}
