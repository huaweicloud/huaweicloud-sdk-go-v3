package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新对象选择信息体
type UpdateDatabaseObjectReq struct {

	// 任务ID
	JobId string `json:"job_id" xml:"job_id"`

	// 是否进行对象选择，是：自定义迁移对象，否：全部迁移，不填默认为否。
	Selected *bool `json:"selected,omitempty" xml:"selected"`

	// 是否库级同步
	SyncDatabase *bool `json:"sync_database,omitempty" xml:"sync_database"`

	// 数据对象选择信息，selected为true时必填。
	Job *[]DatabaseInfo `json:"job,omitempty" xml:"job"`
}

func (o UpdateDatabaseObjectReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabaseObjectReq struct{}"
	}

	return strings.Join([]string{"UpdateDatabaseObjectReq", string(data)}, " ")
}
