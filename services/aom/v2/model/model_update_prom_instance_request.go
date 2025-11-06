package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePromInstanceRequest Request Object
type UpdatePromInstanceRequest struct {

	// 企业项目id。  - 更新单个企业项目下实例，填写企业项目id。
	EnterpriseProjectId string `json:"Enterprise-Project-Id"`

	Body *UpdatePromInstanceRequestModle `json:"body,omitempty"`
}

func (o UpdatePromInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePromInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpdatePromInstanceRequest", string(data)}, " ")
}
