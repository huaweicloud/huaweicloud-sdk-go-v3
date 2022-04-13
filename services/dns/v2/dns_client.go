package v2

import (
	http_client "code.byted.org/ti/huaweicloud-sdk-go-v3/core"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/services/dns/v2/model"
)

type DnsClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDnsClient(hcClient *http_client.HcHttpClient) *DnsClient {
	return &DnsClient{HcClient: hcClient}
}

func DnsClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//创建单个自定义线路
func (c *DnsClient) CreateCustomLine(request *model.CreateCustomLineRequest) (*model.CreateCustomLineResponse, error) {
	requestDef := GenReqDefForCreateCustomLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateCustomLineResponse), nil
	}
}

//删除单个自定义线路
func (c *DnsClient) DeleteCustomLine(request *model.DeleteCustomLineRequest) (*model.DeleteCustomLineResponse, error) {
	requestDef := GenReqDefForDeleteCustomLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteCustomLineResponse), nil
	}
}

//查询所有的云解析服务API版本号列表
func (c *DnsClient) ListApiVersions(request *model.ListApiVersionsRequest) (*model.ListApiVersionsResponse, error) {
	requestDef := GenReqDefForListApiVersions()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListApiVersionsResponse), nil
	}
}

//查询自定义线路
func (c *DnsClient) ListCustomLine(request *model.ListCustomLineRequest) (*model.ListCustomLineResponse, error) {
	requestDef := GenReqDefForListCustomLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListCustomLineResponse), nil
	}
}

//查询名称服务器列表
func (c *DnsClient) ListNameServers(request *model.ListNameServersRequest) (*model.ListNameServersResponse, error) {
	requestDef := GenReqDefForListNameServers()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListNameServersResponse), nil
	}
}

//查询指定的云解析服务API版本号
func (c *DnsClient) ShowApiInfo(request *model.ShowApiInfoRequest) (*model.ShowApiInfoResponse, error) {
	requestDef := GenReqDefForShowApiInfo()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowApiInfoResponse), nil
	}
}

//查询单租户在DNS服务下的资源配额，包括公网zone配额、内网zone配额、Record Set配额、PTR Record配额、入站终端节点配额、出站终端节点配额、自定义线路配额、线路分组配额等。
func (c *DnsClient) ShowDomainQuota(request *model.ShowDomainQuotaRequest) (*model.ShowDomainQuotaResponse, error) {
	requestDef := GenReqDefForShowDomainQuota()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDomainQuotaResponse), nil
	}
}

//更新单个自定义线路
func (c *DnsClient) UpdateCustomLine(request *model.UpdateCustomLineRequest) (*model.UpdateCustomLineResponse, error) {
	requestDef := GenReqDefForUpdateCustomLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateCustomLineResponse), nil
	}
}

//设置弹性IP的PTR记录
func (c *DnsClient) CreateEipRecordSet(request *model.CreateEipRecordSetRequest) (*model.CreateEipRecordSetResponse, error) {
	requestDef := GenReqDefForCreateEipRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateEipRecordSetResponse), nil
	}
}

//查询租户弹性IP的PTR记录列表
func (c *DnsClient) ListPtrRecords(request *model.ListPtrRecordsRequest) (*model.ListPtrRecordsResponse, error) {
	requestDef := GenReqDefForListPtrRecords()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPtrRecordsResponse), nil
	}
}

//将弹性IP的PTR记录恢复为默认值
func (c *DnsClient) RestorePtrRecord(request *model.RestorePtrRecordRequest) (*model.RestorePtrRecordResponse, error) {
	requestDef := GenReqDefForRestorePtrRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.RestorePtrRecordResponse), nil
	}
}

//查询单个弹性IP的PTR记录
func (c *DnsClient) ShowPtrRecordSet(request *model.ShowPtrRecordSetRequest) (*model.ShowPtrRecordSetResponse, error) {
	requestDef := GenReqDefForShowPtrRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPtrRecordSetResponse), nil
	}
}

//修改弹性IP的PTR记录
func (c *DnsClient) UpdatePtrRecord(request *model.UpdatePtrRecordRequest) (*model.UpdatePtrRecordResponse, error) {
	requestDef := GenReqDefForUpdatePtrRecord()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePtrRecordResponse), nil
	}
}

//创建单个Record Set
func (c *DnsClient) CreateRecordSet(request *model.CreateRecordSetRequest) (*model.CreateRecordSetResponse, error) {
	requestDef := GenReqDefForCreateRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRecordSetResponse), nil
	}
}

//创建单个Record Set，仅适用于公网DNS
func (c *DnsClient) CreateRecordSetWithLine(request *model.CreateRecordSetWithLineRequest) (*model.CreateRecordSetWithLineResponse, error) {
	requestDef := GenReqDefForCreateRecordSetWithLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateRecordSetWithLineResponse), nil
	}
}

//删除单个Record Set
func (c *DnsClient) DeleteRecordSet(request *model.DeleteRecordSetRequest) (*model.DeleteRecordSetResponse, error) {
	requestDef := GenReqDefForDeleteRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRecordSetResponse), nil
	}
}

//删除单个Record Set
func (c *DnsClient) DeleteRecordSets(request *model.DeleteRecordSetsRequest) (*model.DeleteRecordSetsResponse, error) {
	requestDef := GenReqDefForDeleteRecordSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteRecordSetsResponse), nil
	}
}

//查询租户Record Set资源列表
func (c *DnsClient) ListRecordSets(request *model.ListRecordSetsRequest) (*model.ListRecordSetsResponse, error) {
	requestDef := GenReqDefForListRecordSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordSetsResponse), nil
	}
}

//查询单个Zone下Record Set列表
func (c *DnsClient) ListRecordSetsByZone(request *model.ListRecordSetsByZoneRequest) (*model.ListRecordSetsByZoneResponse, error) {
	requestDef := GenReqDefForListRecordSetsByZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordSetsByZoneResponse), nil
	}
}

//查询租户Record Set资源列表
func (c *DnsClient) ListRecordSetsWithLine(request *model.ListRecordSetsWithLineRequest) (*model.ListRecordSetsWithLineResponse, error) {
	requestDef := GenReqDefForListRecordSetsWithLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListRecordSetsWithLineResponse), nil
	}
}

//设置Record Set状态
func (c *DnsClient) SetRecordSetsStatus(request *model.SetRecordSetsStatusRequest) (*model.SetRecordSetsStatusResponse, error) {
	requestDef := GenReqDefForSetRecordSetsStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.SetRecordSetsStatusResponse), nil
	}
}

//查询单个Record Set
func (c *DnsClient) ShowRecordSet(request *model.ShowRecordSetRequest) (*model.ShowRecordSetResponse, error) {
	requestDef := GenReqDefForShowRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordSetResponse), nil
	}
}

//查询单个Zone下Record Set列表
func (c *DnsClient) ShowRecordSetByZone(request *model.ShowRecordSetByZoneRequest) (*model.ShowRecordSetByZoneResponse, error) {
	requestDef := GenReqDefForShowRecordSetByZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordSetByZoneResponse), nil
	}
}

//查询单个Record Set，仅适用于公网DNS
func (c *DnsClient) ShowRecordSetWithLine(request *model.ShowRecordSetWithLineRequest) (*model.ShowRecordSetWithLineResponse, error) {
	requestDef := GenReqDefForShowRecordSetWithLine()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowRecordSetWithLineResponse), nil
	}
}

//修改单个Record Set
func (c *DnsClient) UpdateRecordSet(request *model.UpdateRecordSetRequest) (*model.UpdateRecordSetResponse, error) {
	requestDef := GenReqDefForUpdateRecordSet()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRecordSetResponse), nil
	}
}

//修改单个Record Set
func (c *DnsClient) UpdateRecordSets(request *model.UpdateRecordSetsRequest) (*model.UpdateRecordSetsResponse, error) {
	requestDef := GenReqDefForUpdateRecordSets()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdateRecordSetsResponse), nil
	}
}

//为指定实例批量添加或删除标签
func (c *DnsClient) BatchCreateTag(request *model.BatchCreateTagRequest) (*model.BatchCreateTagResponse, error) {
	requestDef := GenReqDefForBatchCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchCreateTagResponse), nil
	}
}

//为指定实例添加标签
func (c *DnsClient) CreateTag(request *model.CreateTagRequest) (*model.CreateTagResponse, error) {
	requestDef := GenReqDefForCreateTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTagResponse), nil
	}
}

//删除资源标签
func (c *DnsClient) DeleteTag(request *model.DeleteTagRequest) (*model.DeleteTagResponse, error) {
	requestDef := GenReqDefForDeleteTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteTagResponse), nil
	}
}

//使用标签查询资源实例
func (c *DnsClient) ListTag(request *model.ListTagRequest) (*model.ListTagResponse, error) {
	requestDef := GenReqDefForListTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagResponse), nil
	}
}

//查询指定实例类型的所有标签集合
func (c *DnsClient) ListTags(request *model.ListTagsRequest) (*model.ListTagsResponse, error) {
	requestDef := GenReqDefForListTags()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListTagsResponse), nil
	}
}

//查询指定实例的标签信息
func (c *DnsClient) ShowResourceTag(request *model.ShowResourceTagRequest) (*model.ShowResourceTagResponse, error) {
	requestDef := GenReqDefForShowResourceTag()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResourceTagResponse), nil
	}
}

//在内网Zone上关联VPC
func (c *DnsClient) AssociateRouter(request *model.AssociateRouterRequest) (*model.AssociateRouterResponse, error) {
	requestDef := GenReqDefForAssociateRouter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AssociateRouterResponse), nil
	}
}

//创建单个内网Zone
func (c *DnsClient) CreatePrivateZone(request *model.CreatePrivateZoneRequest) (*model.CreatePrivateZoneResponse, error) {
	requestDef := GenReqDefForCreatePrivateZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePrivateZoneResponse), nil
	}
}

//创建单个公网Zone
func (c *DnsClient) CreatePublicZone(request *model.CreatePublicZoneRequest) (*model.CreatePublicZoneResponse, error) {
	requestDef := GenReqDefForCreatePublicZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreatePublicZoneResponse), nil
	}
}

//删除单个内网Zone
func (c *DnsClient) DeletePrivateZone(request *model.DeletePrivateZoneRequest) (*model.DeletePrivateZoneResponse, error) {
	requestDef := GenReqDefForDeletePrivateZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePrivateZoneResponse), nil
	}
}

//删除单个公网Zone
func (c *DnsClient) DeletePublicZone(request *model.DeletePublicZoneRequest) (*model.DeletePublicZoneResponse, error) {
	requestDef := GenReqDefForDeletePublicZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeletePublicZoneResponse), nil
	}
}

//在Private Zone上解关联VPC
func (c *DnsClient) DisassociateRouter(request *model.DisassociateRouterRequest) (*model.DisassociateRouterResponse, error) {
	requestDef := GenReqDefForDisassociateRouter()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DisassociateRouterResponse), nil
	}
}

//查询内网Zone的列表
func (c *DnsClient) ListPrivateZones(request *model.ListPrivateZonesRequest) (*model.ListPrivateZonesResponse, error) {
	requestDef := GenReqDefForListPrivateZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPrivateZonesResponse), nil
	}
}

//查询公网Zone的列表
func (c *DnsClient) ListPublicZones(request *model.ListPublicZonesRequest) (*model.ListPublicZonesResponse, error) {
	requestDef := GenReqDefForListPublicZones()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListPublicZonesResponse), nil
	}
}

//查询单个内网Zone
func (c *DnsClient) ShowPrivateZone(request *model.ShowPrivateZoneRequest) (*model.ShowPrivateZoneResponse, error) {
	requestDef := GenReqDefForShowPrivateZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateZoneResponse), nil
	}
}

//查询内网Zone的名称服务器
func (c *DnsClient) ShowPrivateZoneNameServer(request *model.ShowPrivateZoneNameServerRequest) (*model.ShowPrivateZoneNameServerResponse, error) {
	requestDef := GenReqDefForShowPrivateZoneNameServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPrivateZoneNameServerResponse), nil
	}
}

//查询单个公网Zone
func (c *DnsClient) ShowPublicZone(request *model.ShowPublicZoneRequest) (*model.ShowPublicZoneResponse, error) {
	requestDef := GenReqDefForShowPublicZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicZoneResponse), nil
	}
}

//查询单个公网Zone的名称服务器
func (c *DnsClient) ShowPublicZoneNameServer(request *model.ShowPublicZoneNameServerRequest) (*model.ShowPublicZoneNameServerResponse, error) {
	requestDef := GenReqDefForShowPublicZoneNameServer()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowPublicZoneNameServerResponse), nil
	}
}

//修改单个Zone
func (c *DnsClient) UpdatePrivateZone(request *model.UpdatePrivateZoneRequest) (*model.UpdatePrivateZoneResponse, error) {
	requestDef := GenReqDefForUpdatePrivateZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePrivateZoneResponse), nil
	}
}

//修改单个Zone
func (c *DnsClient) UpdatePublicZone(request *model.UpdatePublicZoneRequest) (*model.UpdatePublicZoneResponse, error) {
	requestDef := GenReqDefForUpdatePublicZone()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePublicZoneResponse), nil
	}
}

//设置单个公网Zone状态，支持暂停、启用Zone
func (c *DnsClient) UpdatePublicZoneStatus(request *model.UpdatePublicZoneStatusRequest) (*model.UpdatePublicZoneStatusResponse, error) {
	requestDef := GenReqDefForUpdatePublicZoneStatus()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.UpdatePublicZoneStatusResponse), nil
	}
}
