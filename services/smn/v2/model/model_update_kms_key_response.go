package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKmsKeyResponse Response Object
type UpdateKmsKeyResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 主题下密钥的唯一标识ID。该ID由SMN自动生成。
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateKmsKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKmsKeyResponse struct{}"
	}

	return strings.Join([]string{"UpdateKmsKeyResponse", string(data)}, " ")
}
