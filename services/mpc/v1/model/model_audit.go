package model

import (
	"encoding/json"

	"strings"
)

type Audit struct {
	// 内容质检位置。  取值如下： - 1: 表示原始片源检查。 - 2：表示转码后片源检查，转码后的片源分辨率必须为720p及以上，且格式需与原始片源一致。
	Position *int32 `json:"position,omitempty"`
	// 转码模板ID索引。此索引必须为TransPresetID中的一个。 index从0开始，表示对应的是第几路的校验，比如一进六出，需要对第4路检查，就输入3。  >- 如对原片源质检，此字段不填。 >- 原始分辨率和转码后的分辨率都必须是1280*720及之上。 >- 若对转码后片源检测，必须和原始分辨率一致。
	Index *int32 `json:"index,omitempty"`
}

func (o Audit) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Audit struct{}"
	}

	return strings.Join([]string{"Audit", string(data)}, " ")
}
