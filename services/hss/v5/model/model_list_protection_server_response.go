package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProtectionServerResponse Response Object
type ListProtectionServerResponse struct {

	// **参数解释**: 总数 **取值范围**: 取值0-65535
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 查询勒索防护服务器列表 **取值范围**: 取值0-65535个ProtectionServerInfo对象
	DataList       *[]ProtectionServerInfo `json:"data_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListProtectionServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectionServerResponse struct{}"
	}

	return strings.Join([]string{"ListProtectionServerResponse", string(data)}, " ")
}
