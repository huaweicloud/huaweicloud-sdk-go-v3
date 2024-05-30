package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEncodeServersRequest Request Object
type ListEncodeServersRequest struct {

	// 偏移量为一个大于等于0整数，表示查询该偏移量后面的所有的资源数，默认值为0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页返回的资源个数。取值范围：1~100（默认值为100），一般设置为10、20、50。
	Limit *int32 `json:"limit,omitempty"`

	// 编码服务类型。 - 0：服务器 - 1：容器
	Type *int32 `json:"type,omitempty"`

	// 状态列表。 - 1：运行中 - 2：异常 - 3：重启中 - 4：冻结 - 5：关机 - 100、1014、0：创建中
	Status *int32 `json:"status,omitempty"`

	// 云手机服务器的唯一标识。
	ServerId *string `json:"server_id,omitempty"`
}

func (o ListEncodeServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEncodeServersRequest struct{}"
	}

	return strings.Join([]string{"ListEncodeServersRequest", string(data)}, " ")
}
