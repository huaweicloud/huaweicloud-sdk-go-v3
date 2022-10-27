package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowImageWatermarkWithImageByAddressResponse struct {

	// 当前项目所在region的id，如：xx-xx-1。
	RegionId *string `json:"region_id,omitempty"`

	// 提取出的水印图片存放地址，当前只支持华为云OBS对象，格式为 **obs://bucket/object** ，其中bucket为和当前项目处于同一区域的OBS桶名称，object为对象全路径名。例如：**obs://hwbucket/hwinfo/hw.png**，其中obs://表示OBS存储，hwbucket为桶名，hwinfo/hw.png为对象全路径名。
	ImageWatermark *string `json:"image_watermark,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowImageWatermarkWithImageByAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageWatermarkWithImageByAddressResponse struct{}"
	}

	return strings.Join([]string{"ShowImageWatermarkWithImageByAddressResponse", string(data)}, " ")
}
