package srv6

import (
	"encoding/binary"
	"fmt"

	"github.com/golang/glog"
	"github.com/sbezverk/gobmp/pkg/internal"
)

// BGPPeerNodeSID defines SRv6 BGP Peer Node SID TLV object
// No RFC yet
type BGPPeerNodeSID struct {
	Flag    uint8
	Weight  uint8
	PeerASN uint32
	PeerID  []byte
}

func (b *BGPPeerNodeSID) String(level ...int) string {
	var s string
	l := 0
	if level != nil {
		l = level[0]
	}
	s += internal.AddLevel(l)
	s += "SRv6 BGP Peer Node SID TLV:" + "\n"

	s += internal.AddLevel(l + 1)
	s += fmt.Sprintf("Flag: %02x\n", b.Flag)
	s += internal.AddLevel(l + 1)
	s += fmt.Sprintf("Weight: %d\n", b.Weight)
	s += internal.AddLevel(l + 1)
	s += fmt.Sprintf("Peer ASN: %d\n", b.PeerASN)
	s += internal.AddLevel(l + 1)
	s += fmt.Sprintf("Peer ID: %s\n", internal.MessageHex(b.PeerID))

	return s
}

// MarshalJSON defines a method to Marshal SRv6 BGP Peer Node SID TLV object into JSON format
func (b *BGPPeerNodeSID) MarshalJSON() ([]byte, error) {
	var jsonData []byte
	jsonData = append(jsonData, '{')
	jsonData = append(jsonData, []byte("\"flag\":")...)
	jsonData = append(jsonData, []byte(fmt.Sprintf("%d,", b.Flag))...)
	jsonData = append(jsonData, []byte("\"weight\":")...)
	jsonData = append(jsonData, []byte(fmt.Sprintf("%d,", b.Weight))...)
	jsonData = append(jsonData, []byte("\"peerASN\":")...)
	jsonData = append(jsonData, []byte(fmt.Sprintf("%d,", b.PeerASN))...)
	jsonData = append(jsonData, []byte("\"peerID\":")...)
	jsonData = append(jsonData, []byte(fmt.Sprintf("%s", internal.RawBytesToJSON(b.PeerID)))...)
	jsonData = append(jsonData, '}')

	return jsonData, nil
}

// UnmarshalSRv6BGPPeerNodeSIDTLV builds SRv6 BGP Peer Node SID TLV object
func UnmarshalSRv6BGPPeerNodeSIDTLV(b []byte) (*BGPPeerNodeSID, error) {
	glog.V(6).Infof("SRv6 BGP Peer Node SID TLV Raw: %s", internal.MessageHex(b))
	bgp := BGPPeerNodeSID{}
	p := 0
	bgp.Flag = b[p]
	p++
	bgp.Weight = b[p]
	// Skip reserved 2 bytes
	p += 2
	bgp.PeerASN = binary.BigEndian.Uint32(b[p : p+4])
	p += 4
	bgp.PeerID = make([]byte, 4)
	copy(bgp.PeerID, b[p:p+4])

	return &bgp, nil
}
