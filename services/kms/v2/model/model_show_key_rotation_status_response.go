package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowKeyRotationStatusResponse struct {

	// 密钥轮换状态，默认为“false”，表示关闭密钥轮换功能。
	KeyRotationEnabled *bool `json:"key_rotation_enabled,omitempty" xml:"key_rotation_enabled"`

	// 轮换周期，取值范围为30~365的整数。 周期范围设置根据密钥使用频率进行，若密钥使用频率高，建议设置为短周期；反之，则设置为长周期。
	RotationInterval *int32 `json:"rotation_interval,omitempty" xml:"rotation_interval"`

	// 上一次密钥轮换时间。时间戳，即从1970年1月1日至该时间的总秒数。
	LastRotationTime *string `json:"last_rotation_time,omitempty" xml:"last_rotation_time"`

	// 密钥轮换次数。
	NumberOfRotations *int32 `json:"number_of_rotations,omitempty" xml:"number_of_rotations"`
	HttpStatusCode    int    `json:"-"`
}

func (o ShowKeyRotationStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKeyRotationStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowKeyRotationStatusResponse", string(data)}, " ")
}
