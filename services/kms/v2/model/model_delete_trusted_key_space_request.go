package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTrustedKeySpaceRequest Request Object
type DeleteTrustedKeySpaceRequest struct {

	// **参数解释：** 密钥空间ID。 **约束限制：** 不涉及 **取值范围：** UUID格式，字符长度36-36。 **默认取值：** 不涉及
	KeyspaceId string `json:"keyspace_id"`
}

func (o DeleteTrustedKeySpaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrustedKeySpaceRequest struct{}"
	}

	return strings.Join([]string{"DeleteTrustedKeySpaceRequest", string(data)}, " ")
}
