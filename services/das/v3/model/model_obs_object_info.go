package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsObjectInfo OBS对象信息
type ObsObjectInfo struct {

	// 对象的名称
	ObjectKey *string `json:"object_key,omitempty"`

	// 对象文件的大小
	ContentLength *int64 `json:"content_length,omitempty"`
}

func (o ObsObjectInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsObjectInfo struct{}"
	}

	return strings.Join([]string{"ObsObjectInfo", string(data)}, " ")
}
