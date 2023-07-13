package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunModifyPictureReq 修改图片信息的请求体
type RunModifyPictureReq struct {

	// 图片URL路径，作为图片库中索引图片的ID。
	Path string `json:"path"`

	// 自定义标签，格式为key:value对，其中： - 标签名（key值）须存在于实例中； - 标签内容（value值）为自定义标签值。
	Tags *interface{} `json:"tags"`
}

func (o RunModifyPictureReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunModifyPictureReq struct{}"
	}

	return strings.Join([]string{"RunModifyPictureReq", string(data)}, " ")
}
