package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowImageWatermarkWithImageByAddressRequestBody struct {

	// 项目所在region的id，如北京一为：cn-north-1。
	RegionId string `json:"region_id"`

	// 待提取图片暗水印的图片地址，当前只支持华为云OBS对象，格式为 **obs://bucket/object** ，其中bucket为和当前项目处于同一区域的OBS桶名称，object为对象全路径名。例如：**obs://hwbucket/hwinfo/hw.png**，其中obs://表示OBS存储，hwbucket为桶名，hwinfo/hw.png为对象全路径名。
	SrcFile string `json:"src_file"`

	// 提取出来的水印图片存放地址，格式要求同src_file。
	ImageWatermark string `json:"image_watermark"`
}

func (o ShowImageWatermarkWithImageByAddressRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageWatermarkWithImageByAddressRequestBody struct{}"
	}

	return strings.Join([]string{"ShowImageWatermarkWithImageByAddressRequestBody", string(data)}, " ")
}
