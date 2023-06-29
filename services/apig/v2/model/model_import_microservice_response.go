package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportMicroserviceResponse Response Object
type ImportMicroserviceResponse struct {

	// vpc通道编号
	VpcChannelId *string `json:"vpc_channel_id,omitempty"`

	// api分组编号
	ApiGroupId *string `json:"api_group_id,omitempty"`

	// 导入的api列表
	Apis           *[]MicroserviceImportApiResp `json:"apis,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ImportMicroserviceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportMicroserviceResponse struct{}"
	}

	return strings.Join([]string{"ImportMicroserviceResponse", string(data)}, " ")
}
