package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateQueueCidrRequestBody 修改队列网段的请求body体。
type UpdateQueueCidrRequestBody struct {

	// 队列虚拟私有云网段。  不同CU规格队列支持的网段范围：  - 4cu:  10.0.0.0/8 ~ 10.255.255.192/26  172.16.0.0/12 ~ 172.31.255.192/26  192.168.0.0/16 ~ 192.168.255.192/26  - 16cu:  10.0.0.0/8 ~ 10.255.255.0/24  172.16.0.0/12 ~ 172.31.255.0/24  192.168.0.0/16 ~ 192.168.255.0/24  - 64cu:  10.0.0.0/8 ~ 10.255.252.0/22  172.16.0.0/12 ~ 172.31.252.0/22  192.168.0.0/16 ~ 192.168.252.0/22  - 128cu:  10.0.0.0/8 ~ 10.255.252.0/21  172.16.0.0/12 ~ 172.31.252.0/21  192.168.0.0/16 ~ 192.168.252.0/21
	CidrInVpc *string `json:"cidr_in_vpc,omitempty"`
}

func (o UpdateQueueCidrRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateQueueCidrRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateQueueCidrRequestBody", string(data)}, " ")
}
