package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatakeyCapsuleResponse Response Object
type CreateDatakeyCapsuleResponse struct {

	// **参数解释：** 密钥ID **取值范围：** 不涉及
	KeyId *string `json:"key_id,omitempty"`

	// **参数解释：** datakey和datakey_cipher响应二选一，如果请求参数中没传递public_key，则返回datakey **取值范围：** 不涉及
	Datakey *string `json:"datakey,omitempty"`

	// **参数解释：** datakey和datakey_cipher响应二选一，如果请求参数中传递了public_key，使用public_key加密datakey后返回datakey_cipher **取值范围：** 不涉及
	DatakeyCipher *string `json:"datakey_cipher,omitempty"`

	// **参数解释：** 密钥胶囊 **取值范围：** 不涉及
	DatakeyCapsule *string `json:"datakey_capsule,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDatakeyCapsuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatakeyCapsuleResponse struct{}"
	}

	return strings.Join([]string{"CreateDatakeyCapsuleResponse", string(data)}, " ")
}
