package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FsDirReq 文件系统内合法的目录全路径
type FsDirReq struct {

	// 文件系统内合法的目录全路径。单层目录长度不允许超过255，全路径长度不允许超过4096。
	Path string `json:"path"`
}

func (o FsDirReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FsDirReq struct{}"
	}

	return strings.Join([]string{"FsDirReq", string(data)}, " ")
}
