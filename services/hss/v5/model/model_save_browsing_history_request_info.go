package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveBrowsingHistoryRequestInfo 用户访问记录请求体消息
type SaveBrowsingHistoryRequestInfo struct {

	// 用户页面访问动作in:进入页面，out:离开页面
	Action string `json:"action"`

	// 路径
	Path string `json:"path"`

	// 访问记录id,仅当type为out时携带
	Id *string `json:"id,omitempty"`
}

func (o SaveBrowsingHistoryRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveBrowsingHistoryRequestInfo struct{}"
	}

	return strings.Join([]string{"SaveBrowsingHistoryRequestInfo", string(data)}, " ")
}
