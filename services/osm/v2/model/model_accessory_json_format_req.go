package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessoryJsonFormatReq struct {

	// 文件名称
	AccessoryName *string `json:"accessory_name,omitempty"`

	// 文件来源
	AccessoryFrom *string `json:"accessory_from,omitempty"`

	// 上传类型
	UploadType *int32 `json:"upload_type,omitempty"`

	// 文件内容，Base64格式
	AccessoryData string `json:"accessory_data"`
}

func (o AccessoryJsonFormatReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessoryJsonFormatReq struct{}"
	}

	return strings.Join([]string{"AccessoryJsonFormatReq", string(data)}, " ")
}
