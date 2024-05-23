package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BigKeysInfoResponseBody struct {

	// 大Key所在的DB。
	DbId *int32 `json:"db_id,omitempty"`

	// 大Key类型。
	KeyType *string `json:"key_type,omitempty"`

	// 大Key名。
	KeyName *string `json:"key_name,omitempty"`

	// 大Key的长度。
	KeySize *int32 `json:"key_size,omitempty"`
}

func (o BigKeysInfoResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BigKeysInfoResponseBody struct{}"
	}

	return strings.Join([]string{"BigKeysInfoResponseBody", string(data)}, " ")
}
