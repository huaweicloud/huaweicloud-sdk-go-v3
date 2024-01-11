package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportBackupReq 导入备份请求体
type ImportBackupReq struct {

	// 备份文件路径
	ImportPath string `json:"import_path"`
}

func (o ImportBackupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportBackupReq struct{}"
	}

	return strings.Join([]string{"ImportBackupReq", string(data)}, " ")
}
