/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListVpcRoutesResponse struct {
	// route对象列表
	Routes *[]VpcRoute `json:"routes,omitempty"`
	// 分页信息
	RoutesLinks    *[]NeutronPageLink `json:"routes_links,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListVpcRoutesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListVpcRoutesResponse", string(data)}, " ")
}
