package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperationLogItem 操作
type OperationLogItem struct {

	// 操作时间
	Time *int64 `json:"time,omitempty"`

	// 操作名称,当前已有的action为CREATE_JOB(创建任务),COMMIT_JOB(提交任务),SYSTEM_AUDIT_PASS(系统审核通过),ADMIN_AUDIT_PASS(管理员审核通过),AUDIT_NOT_PASS(审核未通过),TRAINING_FINISH(训练完成),UPLOADING_MODEL(上传语音模型),COMPLETE_JOB(任务完成)
	Action *string `json:"action,omitempty"`

	// 操作者,USER(用户),ADMIN(管理员),SYSTEM(用户)
	Operator *string `json:"operator,omitempty"`

	ExternalInfo *OpExternalInfo `json:"external_info,omitempty"`
}

func (o OperationLogItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationLogItem struct{}"
	}

	return strings.Join([]string{"OperationLogItem", string(data)}, " ")
}
