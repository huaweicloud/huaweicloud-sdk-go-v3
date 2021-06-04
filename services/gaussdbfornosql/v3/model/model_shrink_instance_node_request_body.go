package model

import (
	"encoding/json"

	"strings"
)

type ShrinkInstanceNodeRequestBody struct {
	// 随机缩减的节点数量。该字段取值为1。说明：如果客户端采用直连节点方式，不推荐使用随机节点缩容。

	Num *int32 `json:"num,omitempty"`
	// 指定缩容节点的ID，且该节点必须支持节点缩容。如果该字段不传指定缩减的节点ID，将根据系统内部策略缩减指定个数的节点。说明： - num与node_list必须有一个字段传值。   - 如果num传值时，取值必须为1。   - 如果node_list传值时，其长度必须为1。   - 如果num与node_list同时传值时，则以node_list的值为主。 - 如果node_list取值为空时，缩容以随机节点缩容进行；node_list字段取值不为空，缩容以指定节点ID进行。 - 节点缩容前，请避免直连节点，以防该节点缩容导致业务中断。

	NodeList *[]string `json:"node_list,omitempty"`
}

func (o ShrinkInstanceNodeRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShrinkInstanceNodeRequestBody struct{}"
	}

	return strings.Join([]string{"ShrinkInstanceNodeRequestBody", string(data)}, " ")
}
