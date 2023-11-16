package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyResponse 策略响应
type PolicyResponse struct {

	// id
	Id string `json:"id"`

	// 防护包id
	PackageId string `json:"package_id"`

	// 防护包名
	PackageName string `json:"package_name"`

	// 策略名
	Name string `json:"name"`

	// 描述
	Description string `json:"description"`

	// 所属region的id
	Region string `json:"region"`

	// 清洗阈值
	CleanThreshold int32 `json:"clean_threshold"`

	// 防护ip数
	NumProtectedIp int32 `json:"num_protected_ip"`
}

func (o PolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyResponse struct{}"
	}

	return strings.Join([]string{"PolicyResponse", string(data)}, " ")
}
