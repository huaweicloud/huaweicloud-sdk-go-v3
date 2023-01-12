package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMetadata2Response struct {

	// 元数据是否加密
	Encrypted *bool `json:"encrypted,omitempty"`

	// 秘钥名称
	MasterKeyName *string `json:"master_key_name,omitempty"`

	// 秘钥id
	MasterKeyId *string `json:"master_key_id,omitempty"`

	GesMetadata    *ShowMetadataRespGesMetadata `json:"ges_metadata,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowMetadata2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetadata2Response struct{}"
	}

	return strings.Join([]string{"ShowMetadata2Response", string(data)}, " ")
}
