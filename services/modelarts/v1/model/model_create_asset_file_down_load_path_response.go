package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssetFileDownLoadPathResponse Response Object
type CreateAssetFileDownLoadPathResponse struct {

	// **参数解释**：资产文件下载临时路径。
	Signedurl      *string `json:"signedurl,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAssetFileDownLoadPathResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssetFileDownLoadPathResponse struct{}"
	}

	return strings.Join([]string{"CreateAssetFileDownLoadPathResponse", string(data)}, " ")
}
