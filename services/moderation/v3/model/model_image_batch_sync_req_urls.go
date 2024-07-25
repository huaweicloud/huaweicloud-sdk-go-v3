package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageBatchSyncReqUrls struct {

	// 图片url，目前支持：公网HTTP/HTTPS URL。
	Url string `json:"url"`

	// 图片唯一标识。同一次请求中不可重复，由大小写英文字母、数字、下划线（_）、中划线（-）组成，不超过30个字符。
	DataId string `json:"data_id"`
}

func (o ImageBatchSyncReqUrls) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageBatchSyncReqUrls struct{}"
	}

	return strings.Join([]string{"ImageBatchSyncReqUrls", string(data)}, " ")
}
