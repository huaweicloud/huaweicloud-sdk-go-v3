package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachModelGroupInfo 关联模型分组信息
type AttachModelGroupInfo struct {

	// 分组id。
	Id *string `json:"id,omitempty"`

	// 分组名称。
	Name *string `json:"name,omitempty"`
}

func (o AttachModelGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachModelGroupInfo struct{}"
	}

	return strings.Join([]string{"AttachModelGroupInfo", string(data)}, " ")
}
