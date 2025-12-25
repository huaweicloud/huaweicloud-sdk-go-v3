package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlaybookcontextRef 运行流程的上下文
type PlaybookcontextRef struct {
	DataObject *DataObjectRefInfo `json:"data_object"`
}

func (o PlaybookcontextRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlaybookcontextRef struct{}"
	}

	return strings.Join([]string{"PlaybookcontextRef", string(data)}, " ")
}
