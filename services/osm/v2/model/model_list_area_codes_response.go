package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAreaCodesResponse struct {
	// 国家码列表

	AreaCodeList   *[]AreaCodeSimpleInfoV2 `json:"area_code_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListAreaCodesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAreaCodesResponse struct{}"
	}

	return strings.Join([]string{"ListAreaCodesResponse", string(data)}, " ")
}
