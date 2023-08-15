package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreatePublicipsV2RequestBody This is a auto create Body Object
type BatchCreatePublicipsV2RequestBody struct {
	Bandwidth *BatchBandwidth `json:"bandwidth"`

	Publicip *BatchPublicIp `json:"publicip"`

	// 批量创建EIP的个数
	PublicipNumber int32 `json:"publicip_number"`

	// 企业项目id
	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o BatchCreatePublicipsV2RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePublicipsV2RequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreatePublicipsV2RequestBody", string(data)}, " ")
}
