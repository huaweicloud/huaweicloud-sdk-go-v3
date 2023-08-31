package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchSetTagsReq struct {

	// 临时文件ID
	FileTempId *int64 `json:"file_temp_id,omitempty"`

	// SIM卡id列表，最多500.sim_card_ids与iccids不能同时为空，sim_card_ids参数为空则根据iccids参数处理
	SimCardIds *[]int64 `json:"sim_card_ids,omitempty"`

	// iccid列表，最多支持传入500个iccid。sim_card_ids与iccids不能同时为空，sim_card_ids参数为空则根据iccids参数处理
	Iccids *[]string `json:"iccids,omitempty"`

	// 绑定的标签id列表，最多10
	TagIds *[]int64 `json:"tag_ids,omitempty"`
}

func (o BatchSetTagsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetTagsReq struct{}"
	}

	return strings.Join([]string{"BatchSetTagsReq", string(data)}, " ")
}
