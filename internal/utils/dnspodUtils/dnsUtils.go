/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2022/04/13 17:39 CST
 * @Desc:
 */

package dnspodUtils

import (
	"errors"
	"github.com/AghostPrj/ddns/internal/global"
	"github.com/AghostPrj/ddns/internal/utils/dnsUtils"
	"github.com/nrdcg/dnspod-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
)

var (
	conn     *dnspod.Client
	domainId string
)

func InitDnsClient() {
	clientParams := dnspod.CommonParams{
		LoginToken: viper.GetString(global.ConfDnspodTokenSecretKey),
		Format:     "json",
	}
	conn = dnspod.NewClient(clientParams)
	domainId = ""
}

func AddDomainRecord(domainType, value *string) error {
	subDomain := viper.GetString(global.ConfSubDomainKey)
	result, _, err := conn.Records.Create(domainId, dnspod.Record{
		Name:  subDomain,
		Line:  "默认",
		Type:  *domainType,
		Value: *value,
	})
	log.WithField("err", err).WithField("result", result).WithField("op", "dnspod create record").Debug()
	return err
}

func UpdateDomainRecord(recordId, recordType, value *string) error {
	subDomain := viper.GetString(global.ConfSubDomainKey)
	result, _, err := conn.Records.Update(domainId, *recordId, dnspod.Record{
		Name:  subDomain,
		Line:  "默认",
		Type:  *recordType,
		Value: *value,
	})
	log.WithField("err", err).WithField("result", result).WithField("op", "dnspod update record").Debug()
	return err
}

func DescribeRecord() (result *dnsUtils.DomainRecordPayload, err error) {
	if domainId == "" {
		getDomainId()
	}

	subDomain := viper.GetString(global.ConfSubDomainKey)

	list, _, err := conn.Records.List(domainId, "")
	if err != nil {
		return nil, err
	}

	result = &dnsUtils.DomainRecordPayload{}

	if len(list) > 0 {
		for _, record := range list {
			if record.Name == subDomain {
				ttl, err := strconv.ParseInt(record.TTL, 10, 64)
				if err != nil {
					return nil, err
				}
				switch record.Type {
				case global.DomainTypeIpv4Direct:
					result.Ipv4 = &dnsUtils.DomainRecordInfo{
						Status:   record.Status == "enable",
						TTL:      ttl,
						RecordId: record.ID,
						Value:    record.Value,
					}
					break
				case global.DomainTypeIpv6Direct:
					result.Ipv6 = &dnsUtils.DomainRecordInfo{
						Status:   record.Status == "enable",
						TTL:      ttl,
						RecordId: record.ID,
						Value:    record.Value,
					}
					break
				default:
					break
				}
			}
		}
	}
	return
}

func getDomainId() (err error) {
	list, _, err := conn.Domains.List()
	if err != nil {
		return
	}

	domain := viper.GetString(global.ConfDomainKey)
	for _, domainInfo := range list {
		if domainInfo.Name == domain {
			domainId = domainInfo.ID.String()
			return
		}
	}
	return errors.New("domain not found")
}
