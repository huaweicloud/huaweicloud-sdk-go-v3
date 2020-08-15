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

//
type CreateVpcRouteRequestBody struct {
	Route *CreateVpcRouteOption `json:"route"`
}

func (o CreateVpcRouteRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateVpcRouteRequestBody", string(data)}, " ")
}
