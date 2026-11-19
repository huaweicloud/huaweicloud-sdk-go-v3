package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTrustedKeySpaceResponse Response Object
type CreateTrustedKeySpaceResponse struct {

	// **参数解释：** 创建的密钥空间ID **取值范围：** 不涉及
	KeyspaceId     *string `json:"keyspace_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTrustedKeySpaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrustedKeySpaceResponse struct{}"
	}

	return strings.Join([]string{"CreateTrustedKeySpaceResponse", string(data)}, " ")
}
