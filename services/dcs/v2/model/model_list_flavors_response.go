/*
 * DCS
 *
 * DCS V2版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListFlavorsResponse struct {
	// 产品规格详情。
	Flavors        *[]FlavorsItems `json:"flavors,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListFlavorsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListFlavorsResponse", string(data)}, " ")
}
