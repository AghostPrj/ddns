/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2022/03/19 03:15 CST
 * @Desc:
 */

package runtime

import (
	"github.com/AghostPrj/ddns/internal/global"
	"github.com/AghostPrj/ddns/internal/utils/ipUtils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

func MainLoop() {
	for {
		localIpAddr, err := ipUtils.GetLocalIpByInterfaceName(viper.GetString(global.ConfUpstreamInterfaceNameKey))
		if err != nil {
			log.WithField("err", err).Error()
		} else {
			log.WithField("data", localIpAddr).WithField("op", "get local addr").Debug()

			record, err := global.DescribeRecordFunction()

			if err != nil {
				log.WithField("op", "describe dns record").WithField("err", err).Error()
			} else {
				ipv6Addr := ""
				if len(localIpAddr.V4) > 0 {
					domainRecordType := global.DomainTypeIpv4Direct
					ipv6Addr = "::ffff:" + localIpAddr.V4[0]
					if record.Ipv4 != nil {
						if record.Ipv4.Value != localIpAddr.V4[0] {
							err := global.UpdateDomainRecordFunction(&record.Ipv4.RecordId, &domainRecordType, &localIpAddr.V4[0])
							if err != nil {
								log.WithField("op", "update ipv4 domain record").WithField("err", err).Error()
							} else {
								log.WithField("op", "update ipv4 domain record").WithField("value", localIpAddr.V4[0]).Info()
							}
						}
					} else {
						err := global.AddDomainRecordFunction(&domainRecordType, &localIpAddr.V4[0])
						if err != nil {
							log.WithField("op", "add ipv4 domain record").WithField("err", err).Error()
						} else {
							log.WithField("op", "add ipv4 domain record").WithField("value", localIpAddr.V4[0]).Info()
						}
					}
				}

				if len(localIpAddr.V6) > 0 {
					ipv6Addr = localIpAddr.V6[0]
				}

				if len(ipv6Addr) > 0 {
					domainRecordType := global.DomainTypeIpv6Direct
					if record.Ipv6 != nil {
						if record.Ipv6.Value != ipv6Addr {
							err := global.UpdateDomainRecordFunction(&record.Ipv6.RecordId, &domainRecordType, &ipv6Addr)
							if err != nil {
								log.WithField("op", "update ipv6 domain record").WithField("err", err).Error()
							} else {
								log.WithField("op", "update ipv6 domain record").WithField("value", ipv6Addr).Info()
							}
						}
					} else {
						err := global.AddDomainRecordFunction(&domainRecordType, &ipv6Addr)
						if err != nil {
							log.WithField("op", "add ipv6 domain record").WithField("err", err).Error()
						} else {
							log.WithField("op", "add ipv6 domain record").WithField("value", ipv6Addr).Info()
						}
					}
				}
			}
		}
		time.Sleep(time.Second * viper.GetDuration(global.ConfAppLoopDelayKey))
	}
}
