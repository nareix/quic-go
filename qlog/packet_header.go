package qlog

import (
	"github.com/francoispqt/gojay"

	"github.com/lucas-clemente/quic-go/internal/protocol"
	"github.com/lucas-clemente/quic-go/internal/wire"
)

func transformHeader(hdr *wire.ExtendedHeader) *packetHeader {
	return &packetHeader{
		PacketNumber:     hdr.PacketNumber,
		PayloadLength:    hdr.Length,
		SrcConnectionID:  hdr.SrcConnectionID,
		DestConnectionID: hdr.DestConnectionID,
		Version:          hdr.Version,
	}
}

type packetHeader struct {
	PacketNumber  protocol.PacketNumber `json:"packet_number,string"`
	PacketSize    protocol.ByteCount    `json:"packet_size,omitempty"`
	PayloadLength protocol.ByteCount    `json:"payload_length,omitempty"`

	Version          protocol.VersionNumber `json:"version,omitempty"`
	SrcConnectionID  protocol.ConnectionID  `json:"scid,string,omitempty"`
	DestConnectionID protocol.ConnectionID  `json:"dcid,string,omitempty"`
}

func (h packetHeader) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("packet_number", toString(int64(h.PacketNumber)))
	if h.PacketSize != 0 {
		enc.Int64Key("packet_size", int64(h.PacketSize))
	}
	if h.PayloadLength != 0 {
		enc.Int64Key("payload_length", int64(h.PayloadLength))
	}
	if h.Version != 0 {
		enc.StringKey("version", versionNumber(h.Version).String())
	}
	if h.SrcConnectionID.Len() > 0 {
		enc.StringKey("scil", toString(int64(h.SrcConnectionID.Len())))
		enc.StringKey("scid", connectionID(h.SrcConnectionID).String())
	}
	if h.DestConnectionID.Len() > 0 {
		enc.StringKey("dcil", toString(int64(h.DestConnectionID.Len())))
		enc.StringKey("dcid", connectionID(h.DestConnectionID).String())
	}
}

func (packetHeader) IsNil() bool { return false }
