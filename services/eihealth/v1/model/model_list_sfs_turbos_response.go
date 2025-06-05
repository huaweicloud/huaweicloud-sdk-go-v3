package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSfsTurbosResponse Response Object
type ListSfsTurbosResponse struct {

	// sfs-turbo资源列表。
	SfsTurbos *[]SfsTurboRsp `json:"sfs_turbos,omitempty"`

	// sfs-turbo资源总数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSfsTurbosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSfsTurbosResponse struct{}"
	}

	return strings.Join([]string{"ListSfsTurbosResponse", string(data)}, " ")
}
