package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadKieResponse Response Object
type DownloadKieResponse struct {
	Metadata *DownloadKieResponseBodyMetadata `json:"metadata,omitempty"`

	// 导出的配置项列表。
	Data           *[]CreateKieReq `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o DownloadKieResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadKieResponse struct{}"
	}

	return strings.Join([]string{"DownloadKieResponse", string(data)}, " ")
}
