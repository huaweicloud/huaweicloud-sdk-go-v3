package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchCreateCaTagsRequest struct {

	// 所需要批量创建标签的CA证书ID。
	CaId string `json:"ca_id"`

	Body *BatchOperateTagRequestBody `json:"body,omitempty"`
}

func (o BatchCreateCaTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateCaTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateCaTagsRequest", string(data)}, " ")
}
