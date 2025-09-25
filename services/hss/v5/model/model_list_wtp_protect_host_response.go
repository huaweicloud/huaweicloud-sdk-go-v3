package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWtpProtectHostResponse Response Object
type ListWtpProtectHostResponse struct {

	// **参数解释**: 网页防篡改防护列表信息 **取值范围**: 最小值0，最大值200
	DataList *[]WtpProtectHostResponseInfo `json:"data_list,omitempty"`

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListWtpProtectHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWtpProtectHostResponse struct{}"
	}

	return strings.Join([]string{"ListWtpProtectHostResponse", string(data)}, " ")
}
