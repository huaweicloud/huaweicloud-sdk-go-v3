package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUploadFilesResponse Response Object
type BatchUploadFilesResponse struct {

	// **参数解释**: 上传成功的文件信息列表 **取值范围**: 最小值0，最大值10
	SuccessFileInfo *[]BatchUploadFilesResponseInfoSuccessFileInfo `json:"success_file_info,omitempty"`

	// **参数解释**: 上传失败的文件信息列表 **取值范围**: 最小值0，最大值10
	FailFileInfo   *[]BatchUploadFilesResponseInfoFailFileInfo `json:"fail_file_info,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o BatchUploadFilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUploadFilesResponse struct{}"
	}

	return strings.Join([]string{"BatchUploadFilesResponse", string(data)}, " ")
}
