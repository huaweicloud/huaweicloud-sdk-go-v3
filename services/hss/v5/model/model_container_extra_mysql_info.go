package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContainerExtraMysqlInfo 沙箱类型，MYSQL、MYSQLCHEAT沙箱专用
type ContainerExtraMysqlInfo struct {

	// 自定义反制路径
	CustomPath *string `json:"custom_path,omitempty"`
}

func (o ContainerExtraMysqlInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerExtraMysqlInfo struct{}"
	}

	return strings.Join([]string{"ContainerExtraMysqlInfo", string(data)}, " ")
}
