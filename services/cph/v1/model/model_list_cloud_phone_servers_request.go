package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudPhoneServersRequest Request Object
type ListCloudPhoneServersRequest struct {

	// 偏移量为一个大于等于0整数，表示查询该偏移量后面的所有的资源数，默认值为0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页返回的资源个数。取值范围：1~100（默认值为100），一般设置为10、20、50。
	Limit *int32 `json:"limit,omitempty"`

	// 云手机服务器名称，支持模糊查询。
	ServerName *string `json:"server_name,omitempty"`

	// 云手机服务器的唯一标识。
	ServerId *string `json:"server_id,omitempty"`

	// 云手机服务器是否为自定义网络标识。 - v1：系统定义网络的云手机服务器 - v2：自定义网络的云手机服务器
	NetworkVersion *string `json:"network_version,omitempty"`

	// 手机规格名称。
	PhoneModelName *string `json:"phone_model_name,omitempty"`

	// 查询的起始时间戳。
	CreateSince *int64 `json:"create_since,omitempty"`

	// 查询的结束时间戳。
	CreateUntil *int64 `json:"create_until,omitempty"`

	// 服务器状态。 - 0、1、3、4：创建中 - 2：异常 - 5：正常 - 8：冻结 - 10：关机 - 11：关机中 - 12：关机失败 - 13：开机中
	Status *int32 `json:"status,omitempty"`
}

func (o ListCloudPhoneServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudPhoneServersRequest struct{}"
	}

	return strings.Join([]string{"ListCloudPhoneServersRequest", string(data)}, " ")
}
