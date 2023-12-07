package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateChannelRequest Request Object
type CreateChannelRequest struct {

	// 创建通道时所使用的企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *ChannelCreateReq `json:"body,omitempty"`
}

func (o CreateChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChannelRequest struct{}"
	}

	return strings.Join([]string{"CreateChannelRequest", string(data)}, " ")
}
