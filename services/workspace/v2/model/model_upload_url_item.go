package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadUrlItem OBS 预签名上传地址项。
type UploadUrlItem struct {

	// 区域标识。
	Region *string `json:"region,omitempty"`

	// OBS 预签名上传地址。
	UploadUrl *string `json:"upload_url,omitempty"`

	// OBS 桶名。
	ObsBucket *string `json:"obs_bucket,omitempty"`

	// OBS 对象键。
	ObsObjectKey *string `json:"obs_object_key,omitempty"`
}

func (o UploadUrlItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadUrlItem struct{}"
	}

	return strings.Join([]string{"UploadUrlItem", string(data)}, " ")
}
