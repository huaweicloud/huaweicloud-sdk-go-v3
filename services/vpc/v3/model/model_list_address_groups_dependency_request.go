package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAddressGroupsDependencyRequest Request Object
type ListAddressGroupsDependencyRequest struct {

	// **参数解释**： 用IP地址组的资源ID过滤。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 用IP地址组所属的企业项目ID过滤。 **取值范围**： 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListAddressGroupsDependencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddressGroupsDependencyRequest struct{}"
	}

	return strings.Join([]string{"ListAddressGroupsDependencyRequest", string(data)}, " ")
}
