package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultValueListAttachmentVo 请求的返回的数据对象
type ResultValueListAttachmentVo struct {

	// 实际的数据类型：单个对象，集合 或 NULL
	Value *[]AttachmentVo `json:"value,omitempty"`
}

func (o ResultValueListAttachmentVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultValueListAttachmentVo struct{}"
	}

	return strings.Join([]string{"ResultValueListAttachmentVo", string(data)}, " ")
}
