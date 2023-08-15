package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestPicLayoutBody 子画面信息。
type RestPicLayoutBody struct {
	RestPicLayout *RestPicLayout `json:"restPicLayout,omitempty"`
}

func (o RestPicLayoutBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestPicLayoutBody struct{}"
	}

	return strings.Join([]string{"RestPicLayoutBody", string(data)}, " ")
}
