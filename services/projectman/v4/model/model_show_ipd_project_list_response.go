package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpdProjectListResponse Response Object
type ShowIpdProjectListResponse struct {

	// **参数解释**： 返回状态。 **取值范围**： - success：响应成功 - error：响应失败
	Status *string `json:"status,omitempty"`

	// **参数解释**： 失败信息。 **取值范围**： 不涉及
	Message *string `json:"message,omitempty"`

	// **参数解释**： 项目信息列表。
	Result         *[]ProjectInfoVo `json:"result,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowIpdProjectListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpdProjectListResponse struct{}"
	}

	return strings.Join([]string{"ShowIpdProjectListResponse", string(data)}, " ")
}
