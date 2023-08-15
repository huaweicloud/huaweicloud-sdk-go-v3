package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatTransitIpOption 创建中转IP的请求体。
type CreatTransitIpOption struct {

	// 当前项目子网的ID。
	VirsubnetId string `json:"virsubnet_id"`

	// 中转IP地址。
	IpAddress *string `json:"ip_address,omitempty"`

	// 企业项目ID。创建中转IP时，关联的企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签
	Tags *[]PrivateTag `json:"tags,omitempty"`
}

func (o CreatTransitIpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatTransitIpOption struct{}"
	}

	return strings.Join([]string{"CreatTransitIpOption", string(data)}, " ")
}
