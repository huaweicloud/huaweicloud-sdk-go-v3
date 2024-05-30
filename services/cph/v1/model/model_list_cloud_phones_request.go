package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudPhonesRequest Request Object
type ListCloudPhonesRequest struct {

	// 偏移量为一个大于等于0整数，表示查询该偏移量后面的所有的资源数，默认值为0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页返回的资源个数。取值范围：1~200（默认值为200），一般设置为10、20、50。
	Limit *int32 `json:"limit,omitempty"`

	// 云手机名称，支持模糊查询。
	PhoneName *string `json:"phone_name,omitempty"`

	// 云手机服务器的唯一标识。
	ServerId *string `json:"server_id,omitempty"`

	// 云手机状态。 - 1：创建中 - 2：运行中 - 3：重置中 - 4：重启中 - 6：冻结 - 7：正在关机 - 8：已关机 - -5：重置失败 - -6：重启失败 - -7：手机异常 - -8：创建失败 - -9：关机失败
	Status *int32 `json:"status,omitempty"`

	// 云手机类型。 - 0：普通云手机
	Type *int32 `json:"type,omitempty"`
}

func (o ListCloudPhonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhonesRequest struct{}"
	}

	return strings.Join([]string{"ListCloudPhonesRequest", string(data)}, " ")
}
