package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateV2NetworkResponse Response Object
type CreateV2NetworkResponse struct {

	// **参数解释**：网络连接的ID。 **取值范围**：长度为1~64的英文、数字、中划线组合。
	Id *string `json:"id,omitempty"`

	// **参数解释**：路由网段范围，标准的IPv4格式。 **取值范围**：允许网段为10.0.0.0~10.255.0.0/8~22,172.16.0.0~172.31.0.0/12~22,192.168.0.0~192.168.0.0/16~22。
	VpcCidr        *string `json:"vpc_cidr,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateV2NetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateV2NetworkResponse struct{}"
	}

	return strings.Join([]string{"CreateV2NetworkResponse", string(data)}, " ")
}
