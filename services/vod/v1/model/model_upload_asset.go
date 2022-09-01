package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UploadAsset struct {

	// 媒资所在url
	Url *string `json:"url,omitempty" xml:"url"`

	// 新创建媒资的媒资id
	AssetId *string `json:"asset_id,omitempty" xml:"asset_id"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 错误描述。
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`
}

func (o UploadAsset) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAsset struct{}"
	}

	return strings.Join([]string{"UploadAsset", string(data)}, " ")
}
