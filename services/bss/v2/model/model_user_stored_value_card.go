package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserStoredValueCard struct {

	// 储值卡ID。
	CardId *string `json:"card_id,omitempty"`

	// 储值卡名称。
	CardName *string `json:"card_name,omitempty"`

	// 状态： 1：可使用 2：已用完
	Status *int32 `json:"status,omitempty"`

	// 储值卡面值。
	FaceValue *string `json:"face_value,omitempty"`

	// 储值卡余额。
	Balance *string `json:"balance,omitempty"`

	// 生效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	EffectiveTime *string `json:"effective_time,omitempty"`

	// 失效时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。
	ExpireTime *string `json:"expire_time,omitempty"`
}

func (o UserStoredValueCard) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserStoredValueCard struct{}"
	}

	return strings.Join([]string{"UserStoredValueCard", string(data)}, " ")
}
