package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowImageWatermarkByAddressRequestBody struct {

	// 项目所在region的id，如北京一为：cn-north-1。
	RegionId string `json:"region_id" xml:"region_id"`

	// 待提取文字暗水印的图片地址，当前只支持华为云OBS，格式为 **obs://bucket/object** ，其中bucket为和当前项目处于同一区域的OBS桶名称，object为对象全路径名。例如：**obs://hwbucket/hwinfo/hw.png**，其中obs://表示OBS存储，hwbucket为桶名，hwinfo/hw.png为对象全路径名。
	SrcFile string `json:"src_file" xml:"src_file"`

	// 指定待提取水印的长度，最小0，最大32.。设置后可以提升水印提取性能。
	MarkLen *int32 `json:"mark_len,omitempty" xml:"mark_len"`
}

func (o ShowImageWatermarkByAddressRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageWatermarkByAddressRequestBody struct{}"
	}

	return strings.Join([]string{"ShowImageWatermarkByAddressRequestBody", string(data)}, " ")
}
