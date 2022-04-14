/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2022/03/19 03:15 CST
 * @Desc:
 */

package aliyunDnsUtils

import (
	"github.com/AghostPrj/ddns/internal/global"
	"github.com/AghostPrj/ddns/internal/utils/dnsUtils"
	aliDns "github.com/alibabacloud-go/alidns-20150109/v2/client"
	aliOpenApi "github.com/alibabacloud-go/darabonba-openapi/client"
	aliTea "github.com/alibabacloud-go/tea/tea"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	conn *aliDns.Client
)

func InitDnsClient() {
	tokenId := viper.GetString(global.ConfAliyunTokenIdKey)
	tokenSecret := viper.GetString(global.ConfAliyunTokenSecretKey)
	clientConfig := &aliOpenApi.Config{
		AccessKeyId:     &tokenId,
		AccessKeySecret: &tokenSecret,
		RegionId:        aliTea.String("alidns.cn-hangzhou.aliyuncs.com"),
	}
	c, err := aliDns.NewClient(clientConfig)
	if err != nil {
		log.WithField("op", "init").
			WithField("err", err).Panic("init aliyun client error")
	}
	conn = c
}

func DescribeRecord() (*dnsUtils.DomainRecordPayload, error) {
	domain := viper.GetString(global.ConfDomainKey)
	subDomain := viper.GetString(global.ConfSubDomainKey)
	req := aliDns.DescribeDomainRecordsRequest{
		DomainName: &domain,
		RRKeyWord:  &subDomain,
	}
	records, err := conn.DescribeDomainRecords(&req)
	if err != nil {
		log.WithField("op", "describe domain record").WithField("err", err).Debug()
		return nil, err
	}
	result := dnsUtils.DomainRecordPayload{}

	if *records.Body.TotalCount >= 0 {
		for _, record := range records.Body.DomainRecords.Record {
			switch *record.Type {
			case global.DomainTypeIpv4Direct:
				result.Ipv4 = &dnsUtils.DomainRecordInfo{
					Status:   *record.Status == "ENABLE",
					TTL:      *record.TTL,
					RecordId: *record.RecordId,
					Value:    *record.Value,
				}
				break
			case global.DomainTypeIpv6Direct:
				result.Ipv6 = &dnsUtils.DomainRecordInfo{
					Status:   *record.Status == "ENABLE",
					TTL:      *record.TTL,
					RecordId: *record.RecordId,
					Value:    *record.Value,
				}
				break
			default:
				break
			}
		}
	}

	return &result, nil
}

func UpdateDomainRecord(recordId, recordType, value *string) error {
	subDomain := viper.GetString(global.ConfSubDomainKey)

	req := aliDns.UpdateDomainRecordRequest{RecordId: recordId, RR: &subDomain, Type: recordType, Value: value}
	_, err := conn.UpdateDomainRecord(&req)
	if err != nil {
		log.WithField("op", "update dns record").WithField("err", err).Debug()
		return err
	}
	return err

}

func AddDomainRecord(domainType, value *string) error {
	subDomain := viper.GetString(global.ConfSubDomainKey)
	domain := viper.GetString(global.ConfDomainKey)

	req := aliDns.AddDomainRecordRequest{
		DomainName: &domain,
		RR:         &subDomain,
		Type:       domainType,
		Value:      value,
	}
	_, err := conn.AddDomainRecord(&req)
	if err != nil {
		log.WithField("op", "add dns record").WithField("err", err).Debug()
		return err
	}
	return err
}
