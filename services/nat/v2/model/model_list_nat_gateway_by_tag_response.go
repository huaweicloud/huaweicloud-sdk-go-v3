package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNatGatewayByTagResponse Response Object
type ListNatGatewayByTagResponse struct {

	// 资源对象列表。请参考表Resource字段数据结构说明。
	Resources *[]PublicResource `json:"resources,omitempty"`

	// 总记录数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListNatGatewayByTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatGatewayByTagResponse struct{}"
	}

	return strings.Join([]string{"ListNatGatewayByTagResponse", string(data)}, " ")
}
