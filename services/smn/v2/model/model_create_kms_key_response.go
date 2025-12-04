package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKmsKeyResponse Response Object
type CreateKmsKeyResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 所使用密钥对应的ID。该ID由SMN生成，是主题下的该密钥的唯一标识ID。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateKmsKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKmsKeyResponse struct{}"
	}

	return strings.Join([]string{"CreateKmsKeyResponse", string(data)}, " ")
}
