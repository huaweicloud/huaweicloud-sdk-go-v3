package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateImageWatermarkByAddressRequestBody struct {

	// 当前项目所在region的id，如北京一为：cn-north-1。
	RegionId string `json:"region_id" xml:"region_id"`

	// 待加暗水印的图片地址，当前只支持华为云OBS文件，格式为 **obs://bucket/object** ，其中bucket为和当前项目处于同一区域的OBS桶名称，object为对象全路径名。例如：**obs://hwbucket/hwinfo/hw.png**，其中obs://表示OBS存储，hwbucket为桶名，hwinfo/hw.png为对象全路径名。
	SrcFile string `json:"src_file" xml:"src_file"`

	// 待嵌入的文字暗水印内容，长度不超过32个字符。当前仅支持数字及英文大小写。与图片暗水印image_watermark二选一。
	BlindWatermark *string `json:"blind_watermark,omitempty" xml:"blind_watermark"`

	// 待嵌入的图片暗水印地址，格式要求同src_file字段，与文字暗水印 blind_watermark 二选一，都填写时，生效image_watermark。
	ImageWatermark *string `json:"image_watermark,omitempty" xml:"image_watermark"`

	// 添加水印后的图片存放的地址，格式要求同src_file字段，不设置时，默认取src_file的值，即添加水印后覆盖原文件。
	DstFile *string `json:"dst_file,omitempty" xml:"dst_file"`
}

func (o CreateImageWatermarkByAddressRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageWatermarkByAddressRequestBody struct{}"
	}

	return strings.Join([]string{"CreateImageWatermarkByAddressRequestBody", string(data)}, " ")
}
