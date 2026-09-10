package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttachShareFilesystemResponseBody200Jobs struct {

	// 云手机服务器的唯一标识ID，云手机服务器相关任务包含此字段。
	ServerId *string `json:"server_id,omitempty"`

	// 任务的唯一标识。
	JobId *string `json:"job_id,omitempty"`
}

func (o AttachShareFilesystemResponseBody200Jobs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachShareFilesystemResponseBody200Jobs struct{}"
	}

	return strings.Join([]string{"AttachShareFilesystemResponseBody200Jobs", string(data)}, " ")
}
