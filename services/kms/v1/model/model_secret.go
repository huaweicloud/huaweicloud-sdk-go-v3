package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 凭据对象。
type Secret struct {
	// 凭据的资源标识符。

	Id *string `json:"id,omitempty"`
	// 凭据名称。

	Name *string `json:"name,omitempty"`
	// 凭据状态，取值如下：  ENABLED：表示启用状态  DISABLED：表示禁用状态  PENDING_DELETE：表示待删除状态  FROZEN：表示冻结状态状态

	State *string `json:"state,omitempty"`
	// 用于加密凭据值的KMS主密钥的ID值。

	KmsKeyId *string `json:"kms_key_id,omitempty"`
	// 凭据的描述信息。

	Description *string `json:"description,omitempty"`
	// 凭据创建时间，时间戳，即从1970年1月1日至该时间的总秒数。

	CreateTime *int64 `json:"create_time,omitempty"`
	// 凭据上次更新时间，时间戳，即从1970年1月1日至该时间的总秒数。

	UpdateTime *int64 `json:"update_time,omitempty"`
	// 凭据计划删除时间，时间戳，即从1970年1月1日至该时间的总秒数。  凭据不在删除计划中时，本项值为null。

	ScheduledDeleteTime *int64 `json:"scheduled_delete_time,omitempty"`
}

func (o Secret) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Secret struct{}"
	}

	return strings.Join([]string{"Secret", string(data)}, " ")
}
