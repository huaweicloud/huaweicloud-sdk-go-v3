package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrustedKeySpaceResponse Response Object
type ListTrustedKeySpaceResponse struct {
	PageInfo *ListAccessPointResponseBodyPageInfo `json:"page_info,omitempty"`

	// **参数解释：** 密钥空间列表 **取值范围：** 不涉及
	Keyspaces      *[]ListTrustedKeySpaceResponseBodyKeyspaces `json:"keyspaces,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o ListTrustedKeySpaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrustedKeySpaceResponse struct{}"
	}

	return strings.Join([]string{"ListTrustedKeySpaceResponse", string(data)}, " ")
}
