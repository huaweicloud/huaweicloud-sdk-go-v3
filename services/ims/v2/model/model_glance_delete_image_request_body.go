package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 请求参数
type GlanceDeleteImageRequestBody struct {
	// 取值为：true和false true：表示删除整机镜像时，同时删除其关联的云服务器备份。 false：表示只删除整机镜像，不删除其关联的云服务器备份。

	DeleteBackup *bool `json:"delete_backup,omitempty"`
}

func (o GlanceDeleteImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceDeleteImageRequestBody struct{}"
	}

	return strings.Join([]string{"GlanceDeleteImageRequestBody", string(data)}, " ")
}
