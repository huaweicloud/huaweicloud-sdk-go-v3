package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppWhitelistPolicyProcessExtendResponse Response Object
type ListAppWhitelistPolicyProcessExtendResponse struct {

	// data list
	DataList       *[]AppWhitelistPolicyProcessExtendResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                            `json:"-"`
}

func (o ListAppWhitelistPolicyProcessExtendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppWhitelistPolicyProcessExtendResponse struct{}"
	}

	return strings.Join([]string{"ListAppWhitelistPolicyProcessExtendResponse", string(data)}, " ")
}
