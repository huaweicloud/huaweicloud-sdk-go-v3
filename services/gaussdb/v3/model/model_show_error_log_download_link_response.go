package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowErrorLogDownloadLinkResponse Response Object
type ShowErrorLogDownloadLinkResponse struct {

	// **参数解释**：  数量。  **取值范围**：  不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**：  错误日志下载链接列表。
	List           *[]RdsErrorLogDownload `json:"list,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowErrorLogDownloadLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowErrorLogDownloadLinkResponse struct{}"
	}

	return strings.Join([]string{"ShowErrorLogDownloadLinkResponse", string(data)}, " ")
}
