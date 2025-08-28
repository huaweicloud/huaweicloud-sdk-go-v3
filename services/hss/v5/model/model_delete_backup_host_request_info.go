package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteBackupHostRequestInfo struct {

	// **参数解释**: 需要删除的备份服务器ID列表，填写的备份服务器ID必须是已运行中的备份服务器ID **约束限制**: 需要使用 ListBackupHostsInfo 接口查询已运行中的远端备份服务器，在 ListBackupHostsInfo 接口的响应体中，backup_host_status 等于 1 的 backup_host_id 是已运行的远端备份服务器ID **取值范围**: 不超过100条 **默认取值**: 不涉及
	BackupHostIdList []string `json:"backup_host_id_list"`
}

func (o DeleteBackupHostRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupHostRequestInfo struct{}"
	}

	return strings.Join([]string{"DeleteBackupHostRequestInfo", string(data)}, " ")
}
