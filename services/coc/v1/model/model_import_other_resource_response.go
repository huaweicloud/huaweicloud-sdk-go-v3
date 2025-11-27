package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportOtherResourceResponse Response Object
type ImportOtherResourceResponse struct {

	// **参数解释：** 成功数即excel表中数据通过校验成功的数据条数。 **取值范围：** 不涉及。
	Success *int32 `json:"success,omitempty"`

	// **参数解释：** 总条数即excel表中数据的总数据条数。 **取值范围：** 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释：** 失败信息集合即excel表格中校验失败报错信息组合。 **取值范围：** 不涉及。
	ErrorList      *[]string `json:"error_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ImportOtherResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportOtherResourceResponse struct{}"
	}

	return strings.Join([]string{"ImportOtherResourceResponse", string(data)}, " ")
}
