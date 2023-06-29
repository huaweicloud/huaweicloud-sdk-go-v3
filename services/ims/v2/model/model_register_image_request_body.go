package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterImageRequestBody 镜像上传请求体
type RegisterImageRequestBody struct {

	// 源镜像的URL，格式：<bucket>:<file> image_url对应的镜像桶中的文件，镜像文件格式的取值范围为：ZVHD、QCOW2、VHD、RAW、VHDX、QED、VDI、QCOW、ZVHD2、VMDK。
	ImageUrl string `json:"image_url"`
}

func (o RegisterImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterImageRequestBody struct{}"
	}

	return strings.Join([]string{"RegisterImageRequestBody", string(data)}, " ")
}
