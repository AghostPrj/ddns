/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2022/03/19 03:15 CST
 * @Desc:
 */

package ipUtils

import (
	"errors"
	"github.com/Asphaltt/go-iproute2/ip"
	log "github.com/sirupsen/logrus"
	"net"
)

var (
	conn *ip.Client
)

func InitNetLinkConn() {
	c, err := ip.New()
	if err != nil {
		log.WithField("op", "init").
			WithField("err", err).Panic()
	}
	conn = c
}

func GetLocalIpByInterfaceName(interfaceName string) (*IpAddrPayload, error) {
	links, err := conn.ListLinks()
	if err != nil {
		log.WithField("op", "list ip link").WithField("err", err).Debug()
		return nil, err
	}

	linkFindFlag := false
	linkId := -1

	for _, link := range links {
		if link.Name == interfaceName {
			linkFindFlag = true
			linkId = link.Ifindex
		}
	}

	if !linkFindFlag {
		log.WithField("op", "list ip link").WithField("err", "interface not found").Debug()
		return nil, errors.New("interface not found")
	}

	addresses, err := conn.ListAddresses()
	if err != nil {
		log.WithField("op", "list ip addr").WithField("err", err).Debug()
		return nil, err
	}

	resultPayload := IpAddrPayload{V4: make([]string, 0), V6: make([]string, 0)}

	for _, entry := range addresses[linkId] {
		var addr net.IP
		if entry.LocalAddr != nil {
			addr = entry.LocalAddr
		} else {
			addr = entry.InterfaceAddr
		}

		if addr.IsLoopback() || addr.IsPrivate() ||
			(len(addr) == 4 && addr[0] == 169 && addr[1] == 254) ||
			(len(addr) == 16 && addr[0] == 0xfe && addr[1] == 0x80) {
			continue
		}
		if addr.To4() != nil {
			log.WithField("addr", addr.String()).WithField("flag", entry.AddrFlags.String()).WithField("ent", entry).Debug()
			resultPayload.V4 = append(resultPayload.V4, addr.String())
		} else {

			log.WithField("addr", addr.String()).WithField("flag", entry.AddrFlags.String()).WithField("ent", entry).Debug()
			resultPayload.V6 = append(resultPayload.V6, addr.String())
		}
	}

	return &resultPayload, nil

}
