package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpBindingBody ip绑定请求体
type IpBindingBody struct {

	// 实例id
	PackageId string `json:"package_id"`

	// 防护ip的id列表
	IdList []string `json:"id_list"`
}

func (o IpBindingBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpBindingBody struct{}"
	}

	return strings.Join([]string{"IpBindingBody", string(data)}, " ")
}
