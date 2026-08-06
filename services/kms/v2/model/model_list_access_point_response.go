package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessPointResponse Response Object
type ListAccessPointResponse struct {
	PageInfo *ListAccessPointResponseBodyPageInfo `json:"page_info,omitempty"`

	// **参数解释：** 接入点列表 **取值范围：** 不涉及
	AccessPoints   *[]ListAccessPointResponseBodyAccessPoints `json:"access_points,omitempty"`
	HttpStatusCode int                                        `json:"-"`
}

func (o ListAccessPointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessPointResponse struct{}"
	}

	return strings.Join([]string{"ListAccessPointResponse", string(data)}, " ")
}
