package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSubCustomersResponse struct {
	// 子用户列表

	SubCustomerInfos *[]SubCutomerInfoV2 `json:"sub_customer_infos,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o ListSubCustomersResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSubCustomersResponse struct{}"
	}

	return strings.Join([]string{"ListSubCustomersResponse", string(data)}, " ")
}
