package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBlockchainDetailResponse struct {
	BasicInfo *BasicInfo `json:"basic_info,omitempty" xml:"basic_info"`

	// 通道信息
	Channels *[]ChannelInfo `json:"channels,omitempty" xml:"channels"`

	// peer组织信息
	PeerInfo *[]PeerInfo `json:"peer_info,omitempty" xml:"peer_info"`

	// light_peer组织信息
	LightPeerInfo *[]PeerInfo `json:"light_peer_info,omitempty" xml:"light_peer_info"`

	OrdererInfo *PeerInfo `json:"orderer_info,omitempty" xml:"orderer_info"`

	CouchDbInfo *CouchDbInfo `json:"couch_db_info,omitempty" xml:"couch_db_info"`

	DmsKafkaInfo *DmsKafkaInfo `json:"dms_kafka_info,omitempty" xml:"dms_kafka_info"`

	IefInfo *IefInfo `json:"ief_info,omitempty" xml:"ief_info"`

	SfsInfo *SfsInfo `json:"sfs_info,omitempty" xml:"sfs_info"`

	AgentInfo *PeerInfo `json:"agent_info,omitempty" xml:"agent_info"`

	RestapiInfo *PeerInfo `json:"restapi_info,omitempty" xml:"restapi_info"`

	// 云硬盘存储卷信息
	EvsPvcInfo map[string]map[string]string `json:"evs_pvc_info,omitempty" xml:"evs_pvc_info"`

	Tc3TaskserverInfo *PeerInfo `json:"tc3_taskserver_info,omitempty" xml:"tc3_taskserver_info"`

	ObsBucketInfo  *ObsInfo `json:"obs_bucket_info,omitempty" xml:"obs_bucket_info"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowBlockchainDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlockchainDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowBlockchainDetailResponse", string(data)}, " ")
}
