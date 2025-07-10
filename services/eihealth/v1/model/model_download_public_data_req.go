package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadPublicDataReq 下载文件请求体
type DownloadPublicDataReq struct {
	Type *DownloadPublicDataTypeEnum `json:"type,omitempty"`

	// **参数解释**：   资源地址，支持https、obs地址。   **约束限制**：   不涉及   **取值范围**：   长度为[1-2000]个字符。 **默认取值**：   不涉及
	Url *string `json:"url,omitempty"`
}

func (o DownloadPublicDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadPublicDataReq struct{}"
	}

	return strings.Join([]string{"DownloadPublicDataReq", string(data)}, " ")
}
