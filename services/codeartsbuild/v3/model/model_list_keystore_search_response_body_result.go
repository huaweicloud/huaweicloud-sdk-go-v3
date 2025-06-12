package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKeystoreSearchResponseBodyResult 租户下文件列表
type ListKeystoreSearchResponseBodyResult struct {

	// 总数
	Total float32 `json:"total,omitempty"`

	// 文件列表
	KeystoreList *[]ListKeystoreSearchResponseBodyResultKeystoreList `json:"keystore_list,omitempty"`
}

func (o ListKeystoreSearchResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeystoreSearchResponseBodyResult struct{}"
	}

	return strings.Join([]string{"ListKeystoreSearchResponseBodyResult", string(data)}, " ")
}
