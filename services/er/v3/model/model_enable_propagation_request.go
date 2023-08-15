package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnablePropagationRequest Request Object
type EnablePropagationRequest struct {

	// 幂等性标识
	XClientToken *string `json:"X-Client-Token,omitempty"`

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 路由表ID
	RouteTableId string `json:"route_table_id"`

	Body *PropagationRequestBody `json:"body,omitempty"`
}

func (o EnablePropagationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnablePropagationRequest struct{}"
	}

	return strings.Join([]string{"EnablePropagationRequest", string(data)}, " ")
}
