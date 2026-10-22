package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPresignedUrlRequest Request Object
type ShowPresignedUrlRequest struct {

	// **参数解释：** 操作名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	OperationName string `json:"operation_name"`
}

func (o ShowPresignedUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPresignedUrlRequest struct{}"
	}

	return strings.Join([]string{"ShowPresignedUrlRequest", string(data)}, " ")
}
