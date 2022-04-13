package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBlockchainDetailResponse struct {
	BasicInfo *BasicInfo `json:"basic_info,omitempty"`
	// 通道信息

	Channels *[]ChannelInfo `json:"channels,omitempty"`
	// peer组织信息

	PeerInfo *[]PeerInfo `json:"peer_info,omitempty"`
	// light_peer组织信息

	LightPeerInfo *[]PeerInfo `json:"light_peer_info,omitempty"`

	OrdererInfo *PeerInfo `json:"orderer_info,omitempty"`

	CouchDbInfo *CouchDbInfo `json:"couch_db_info,omitempty"`

	DmsKafkaInfo *DmsKafkaInfo `json:"dms_kafka_info,omitempty"`

	IefInfo *IefInfo `json:"ief_info,omitempty"`

	SfsInfo *SfsInfo `json:"sfs_info,omitempty"`

	AgentInfo *PeerInfo `json:"agent_info,omitempty"`

	RestapiInfo *PeerInfo `json:"restapi_info,omitempty"`
	// 云硬盘存储卷信息

	EvsPvcInfo map[string]map[string]string `json:"evs_pvc_info,omitempty"`

	Tc3TaskserverInfo *PeerInfo `json:"tc3_taskserver_info,omitempty"`

	ObsBucketInfo  *ObsInfo `json:"obs_bucket_info,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowBlockchainDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlockchainDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowBlockchainDetailResponse", string(data)}, " ")
}
