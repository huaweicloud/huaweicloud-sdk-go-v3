package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeCloudServiceResponse Response Object
type SubscribeCloudServiceResponse struct {

	// 响应状态码
	Code *int32 `json:"code,omitempty"`

	// 响应信息
	Message *string `json:"message,omitempty"`

	// 包年/包月订单ID,按需场景为空。
	OrderIds *[]string `json:"order_ids,omitempty"`

	// jobIds,包年/包月场景为空。
	JobIds *[]string `json:"job_ids,omitempty"`

	// jobId,包年/包月场景为空。
	JobId *string `json:"job_id,omitempty"`

	// 产品名称
	ProductName    *string `json:"product_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SubscribeCloudServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeCloudServiceResponse struct{}"
	}

	return strings.Join([]string{"SubscribeCloudServiceResponse", string(data)}, " ")
}
