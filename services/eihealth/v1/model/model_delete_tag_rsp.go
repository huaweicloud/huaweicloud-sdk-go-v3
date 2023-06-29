package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTagRsp 批量删除镜像tag返回体
type DeleteTagRsp struct {

	// 镜像tag名称
	Tag *string `json:"tag,omitempty"`

	// 删除结果
	Status *string `json:"status,omitempty"`

	// 失败原因
	FailedReason *string `json:"failed_reason,omitempty"`
}

func (o DeleteTagRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagRsp struct{}"
	}

	return strings.Join([]string{"DeleteTagRsp", string(data)}, " ")
}
