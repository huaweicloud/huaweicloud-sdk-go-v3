package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogBackupResponse Response Object
type ShowLogBackupResponse struct {
	LogList *[]LogList `json:"logList,omitempty"`

	// 查询日志的类型。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 日志文件是否已经查询完。 **取值范围**： - true： 日志文件已经查询完，没有更多结果了。 - false：日志文件未查询完，因日志条数已达到请求条数或者日志大小达到1MB，查询提前返回结果。
	Completed      *bool `json:"completed,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowLogBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogBackupResponse struct{}"
	}

	return strings.Join([]string{"ShowLogBackupResponse", string(data)}, " ")
}
