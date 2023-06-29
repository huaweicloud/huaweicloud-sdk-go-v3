package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCaTagRequest Request Object
type CreateCaTagRequest struct {

	// 所需要创建标签的CA证书ID。
	CaId string `json:"ca_id"`

	Body *ResourceTagRequestBody `json:"body,omitempty"`
}

func (o CreateCaTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCaTagRequest struct{}"
	}

	return strings.Join([]string{"CreateCaTagRequest", string(data)}, " ")
}
