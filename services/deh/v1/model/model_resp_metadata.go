package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 弹性云服务器元数据。
type RespMetadata struct {

	// 弹性云服务器系统类型。
	OsType *string `json:"os_type,omitempty" xml:"os_type"`
}

func (o RespMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RespMetadata struct{}"
	}

	return strings.Join([]string{"RespMetadata", string(data)}, " ")
}
