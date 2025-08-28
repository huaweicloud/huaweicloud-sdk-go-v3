package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackupHostsInfoResponse Response Object
type ListBackupHostsInfoResponse struct {

	// **参数解释**: 远端备份服务器列表 **取值范围**: 最小值0，最大值200
	DataList *[]WtpBackupHostResponseInfo `json:"data_list,omitempty"`

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBackupHostsInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupHostsInfoResponse struct{}"
	}

	return strings.Join([]string{"ListBackupHostsInfoResponse", string(data)}, " ")
}
