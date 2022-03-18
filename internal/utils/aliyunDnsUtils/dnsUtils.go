/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2022/03/19 03:15 CST
 * @Desc:
 */

package aliyunDnsUtils

import (
	"github.com/AghostPrj/ddns/internal/global"
	aliDns "github.com/alibabacloud-go/alidns-20150109/v2/client"
	aliOpenApi "github.com/alibabacloud-go/darabonba-openapi/client"
	aliTea "github.com/alibabacloud-go/tea/tea"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	conn *aliDns.Client
)

func InitAliyunDnsClient() {
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

func DescribeRecord() {
	domain := viper.GetString(global.ConfDomainKey)
	subDomain := viper.GetString(global.ConfSubDomainKey)
	req := aliDns.DescribeDomainRecordsRequest{
		DomainName: &domain,
		RRKeyWord:  &subDomain,
	}
	conn.DescribeDomainRecords()
}
