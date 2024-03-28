package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterAuthorizedQueueRequestBody 队列赋权的请求参数。
type RegisterAuthorizedQueueRequestBody struct {

	// 队列名称。
	QueueName string `json:"queue_name"`

	// 被赋权用户名称。给该用户赋使用队列的权限，回收其使用权限，或者更新其使用权限。
	UserName string `json:"user_name"`

	// 指定赋权或回收。值为：grant，revoke或update。当用户同时拥有grant和revoke权限的时候才有权限使用update操作。 grant：赋权。 revoke：回收权限。 update：清空原来的所有权限，赋予本次提供的权限数组中的权限。
	Action string `json:"action"`

	// 待赋权、回收或更新的权限列表。可操作的权限可以是以下三种权限中的一种或多种。 SUBMIT_JOB：提交作业 CANCEL_JOB ：取消作业 DROP_QUEUE ：删除队列 说明： 若需更新的权限列表为空，则表示回收用户在该队列的所有权限。
	Privileges []string `json:"privileges"`
}

func (o RegisterAuthorizedQueueRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterAuthorizedQueueRequestBody struct{}"
	}

	return strings.Join([]string{"RegisterAuthorizedQueueRequestBody", string(data)}, " ")
}
