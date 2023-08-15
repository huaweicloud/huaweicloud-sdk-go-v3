package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImageWatermarkByAddressResponse Response Object
type CreateImageWatermarkByAddressResponse struct {

	// 当前项目所在region的id，如：xx-xx-1。
	RegionId *string `json:"region_id,omitempty"`

	// 添加水印后的OBS图片地址，当前只支持华为云OBS文件，格式为 **obs://bucket/object** ，其中bucket为和当前项目处于同一区域的OBS桶名称，object为对象全路径名。例如：**obs://hwbucket/hwinfo/hw.png**，其中obs://表示OBS存储，hwbucket为桶名，hwinfo/hw.png为对象全路径名。
	WatermarkedFile *string `json:"watermarked_file,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o CreateImageWatermarkByAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageWatermarkByAddressResponse struct{}"
	}

	return strings.Join([]string{"CreateImageWatermarkByAddressResponse", string(data)}, " ")
}
