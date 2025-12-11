package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFilePathWhiteListsResponse Response Object
type ShowFilePathWhiteListsResponse struct {

	// **参数解释**: 默认路径，不可编辑 **取值范围**: 最小值0，最大值20000
	DefaultPath *[]string `json:"default_path,omitempty"`

	// **参数解释**: 自定义路径，选填，可编辑 **取值范围**: 最小值0，最大值20000
	CustomizedPath *[]string `json:"customized_path,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowFilePathWhiteListsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFilePathWhiteListsResponse struct{}"
	}

	return strings.Join([]string{"ShowFilePathWhiteListsResponse", string(data)}, " ")
}
