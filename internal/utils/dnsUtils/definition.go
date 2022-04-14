/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2022/03/19 05:46 CST
 * @Desc:
 */

package dnsUtils

type DomainRecordPayload struct {
	Ipv4 *DomainRecordInfo
	Ipv6 *DomainRecordInfo
}

type DomainRecordInfo struct {
	Status   bool
	TTL      int64
	RecordId string
	Value    string
}
