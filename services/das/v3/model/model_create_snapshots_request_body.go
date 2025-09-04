package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSnapshotsRequestBody struct {

	// 需要创建的snapshot类型，0表示元数据锁，1表示innodb等待锁，2表示最近发生的死锁
	Modules []int32 `json:"modules"`
}

func (o CreateSnapshotsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSnapshotsRequestBody", string(data)}, " ")
}
