package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MultipartUploadInfo 训练视频已上传分片信息
type MultipartUploadInfo struct {

	// 分片编号
	PartNumber *string `json:"part_number,omitempty"`

	// 分片文件标识
	Etag *string `json:"etag,omitempty"`
}

func (o MultipartUploadInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultipartUploadInfo struct{}"
	}

	return strings.Join([]string{"MultipartUploadInfo", string(data)}, " ")
}
