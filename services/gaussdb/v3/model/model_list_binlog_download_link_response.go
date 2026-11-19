package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBinlogDownloadLinkResponse Response Object
type ListBinlogDownloadLinkResponse struct {

	// **参数解释**：  Binlog日志下载链接列表，有效时间为5分钟。
	Links          *[]string `json:"links,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListBinlogDownloadLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBinlogDownloadLinkResponse struct{}"
	}

	return strings.Join([]string{"ListBinlogDownloadLinkResponse", string(data)}, " ")
}
