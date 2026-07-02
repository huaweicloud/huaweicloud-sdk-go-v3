package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DownLoadFileInfoItem struct {

	// **参数解释**：  日志文件ID。  **取值范围**：  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：  日志文件名称。  **取值范围**：  不涉及。
	FileName *string `json:"file_name,omitempty"`

	// **参数解释**：  日志文件大小，单位为字节。  **取值范围**：  不涉及。
	FileSize *int32 `json:"file_size,omitempty"`

	// **参数解释**：  日志下载链接。  **取值范围**：  不涉及。
	DownloadUrl *string `json:"download_url,omitempty"`

	// **参数解释**：  下载链接过期时间，格式为\"yyyy-MM-dd HH:mm:ss\"。  **取值范围**：  不涉及。
	ExpireTime *string `json:"expire_time,omitempty"`
}

func (o DownLoadFileInfoItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownLoadFileInfoItem struct{}"
	}

	return strings.Join([]string{"DownLoadFileInfoItem", string(data)}, " ")
}
