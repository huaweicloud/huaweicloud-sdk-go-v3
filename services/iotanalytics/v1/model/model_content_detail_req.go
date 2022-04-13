package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IoTA服务各类数据源详细配置内容
type ContentDetailReq struct {
	IotdaContent *IotdaContentReq `json:"iotda_content,omitempty"`

	ObsContent *ObsContentReq `json:"obs_content,omitempty"`

	DisContent *DisContentReq `json:"dis_content,omitempty"`

	SmnContent *SmnContentReq `json:"smn_content,omitempty"`

	FunctionGraphContent *FunctionGraphContentReq `json:"function_graph_content,omitempty"`

	ModelArtsContent *ModelArtsContentReq `json:"model_arts_content,omitempty"`

	DcsContent *DcsContentReq `json:"dcs_content,omitempty"`

	KafkaContent *KafkaContentReq `json:"kafka_content,omitempty"`

	ApiContent *ApiContentReq `json:"api_content,omitempty"`

	NodeContent *NodeContentReq `json:"node_content,omitempty"`

	EdgeContent *EdgeContentReq `json:"edge_content,omitempty"`
}

func (o ContentDetailReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContentDetailReq struct{}"
	}

	return strings.Join([]string{"ContentDetailReq", string(data)}, " ")
}
