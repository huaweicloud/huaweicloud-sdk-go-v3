package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDirectoryResultData data，统一的返回结果的最外层数据结构。
type CreateDirectoryResultData struct {
	Value *DirectoryVo `json:"value,omitempty"`
}

func (o CreateDirectoryResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDirectoryResultData struct{}"
	}

	return strings.Join([]string{"CreateDirectoryResultData", string(data)}, " ")
}
