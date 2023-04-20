package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateRedislogDownloadLinkResponse struct {

	// 日志id
	Id *string `json:"id,omitempty"`

	// 后台任务id
	BackupId *string `json:"backup_id,omitempty"`

	// 日志下载链接，默认有效时间为24小时
	Link           *string `json:"link,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRedislogDownloadLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRedislogDownloadLinkResponse struct{}"
	}

	return strings.Join([]string{"CreateRedislogDownloadLinkResponse", string(data)}, " ")
}
