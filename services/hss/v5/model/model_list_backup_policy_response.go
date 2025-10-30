package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackupPolicyResponse Response Object
type ListBackupPolicyResponse struct {

	// **参数解释** 总数 **取值范围** 取值0-2097152
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** 备份策略列表 **取值范围** 取值0-10241
	DataList       *[]GetBackupPolicyResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListBackupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListBackupPolicyResponse", string(data)}, " ")
}
