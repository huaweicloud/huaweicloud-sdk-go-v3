package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObsObject struct {

	// 对象名
	FileName string `json:"file_name"`

	// 文件大小，单位KB
	Size int32 `json:"size"`

	// 上次修改时间，格式如：2020/07/16 15:11:55 GMT+08:00
	LastModified string `json:"last_modified"`
}

func (o ObsObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsObject struct{}"
	}

	return strings.Join([]string{"ObsObject", string(data)}, " ")
}
