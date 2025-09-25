package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteConfReqNew struct {

	// 配置文件名称。
	Name string `json:"name"`
}

func (o DeleteConfReqNew) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConfReqNew struct{}"
	}

	return strings.Join([]string{"DeleteConfReqNew", string(data)}, " ")
}
