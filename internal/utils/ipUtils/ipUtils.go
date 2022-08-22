/**
 * @Author: aghost
 * @Author: ggg17226@gmail.com
 * @Date: 2022/03/19 03:15 CST
 * @Desc:
 */

package ipUtils

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Asphaltt/go-iproute2/ip"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"net/http"
	"time"
)

type AddrServiceResultPayload struct {
	SourceIp string `json:"source_ip"`
}

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

func GetLocalIpByRemoteService() (*IpAddrPayload, error) {
	ip, err := getLocalV4IpByRemoteService()
	if err != nil {
		return nil, err
	}

	result := IpAddrPayload{}
	result.V4 = make([]string, 1)
	result.V4[0] = ip
	result.V6 = make([]string, 0)

	return &result, nil
}

func getLocalV4IpByRemoteService() (string, error) {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		dialer := &net.Dialer{
			Timeout:   15 * time.Second,
			KeepAlive: 60 * time.Second,
		}
		return dialer.DialContext(ctx, "tcp4", addr)
	}
	httpClient := &http.Client{
		Transport: transport,
		Timeout:   15 * time.Second,
	}

	resp, err := httpClient.Get("https://addr.agh0st.com")

	if err != nil {
		return "", err
	}

	if resp.Body == nil {
		return "", errors.New("body empty")
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	payload := AddrServiceResultPayload{}

	err = json.Unmarshal(bodyBytes, &payload)
	if err != nil {
		return "", err
	}

	return payload.SourceIp, nil
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
