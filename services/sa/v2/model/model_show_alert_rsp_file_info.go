package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowAlertRspFileInfo struct {

	// The name, display only
	FilePath *string `json:"file_path,omitempty"`

	// The name, display only
	FileContent *string `json:"file_content,omitempty"`

	// The name, display only
	FileNewPath *string `json:"file_new_path,omitempty"`

	// The name, display only
	FileHash *string `json:"file_hash,omitempty"`

	// The name, display only
	FileMd5 *string `json:"file_md5,omitempty"`

	// The name, display only
	FileSha256 *string `json:"file_sha256,omitempty"`

	// The name, display only
	FileAttr *string `json:"file_attr,omitempty"`
}

func (o ShowAlertRspFileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertRspFileInfo struct{}"
	}

	return strings.Join([]string{"ShowAlertRspFileInfo", string(data)}, " ")
}
