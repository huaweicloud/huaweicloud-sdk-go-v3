package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessKeysV5Response Response Object
type ListAccessKeysV5Response struct {

	// 永久访问密钥列表。
	AccessKeys *[]AccessKeyMetadata `json:"access_keys,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAccessKeysV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessKeysV5Response struct{}"
	}

	return strings.Join([]string{"ListAccessKeysV5Response", string(data)}, " ")
}
