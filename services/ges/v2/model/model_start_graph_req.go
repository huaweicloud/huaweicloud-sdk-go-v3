package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartGraphReq 启动图请求体
type StartGraphReq struct {

	// 启动图时关联的备份ID，设置此参数时，表示从备份进行启动；如果为空，表示从上次关闭图时的状态启动。
	GraphBackupId *string `json:"graph_backup_id,omitempty"`
}

func (o StartGraphReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartGraphReq struct{}"
	}

	return strings.Join([]string{"StartGraphReq", string(data)}, " ")
}
