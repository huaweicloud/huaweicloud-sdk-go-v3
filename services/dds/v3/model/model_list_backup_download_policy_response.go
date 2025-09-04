package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackupDownloadPolicyResponse Response Object
type ListBackupDownloadPolicyResponse struct {

	// 备份下载策略id。
	Id *string `json:"id,omitempty"`

	// 备份下载开关。open表示打开备份下载开关，允许外网下载。close表示关闭备份下载开关，不允许外网下载。
	Action         *string `json:"action,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListBackupDownloadPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupDownloadPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListBackupDownloadPolicyResponse", string(data)}, " ")
}
