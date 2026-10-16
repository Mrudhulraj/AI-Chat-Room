package webrtc

import (
	"live-video-meet/pkg/chat"
	"sync"

	"github.com/pion/webrtc/v3"
)

type Peers struct {
	ListLock    sync.RWMutex
	Connections []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

type Room struct {
	Peers *Peers
	hub   *chat.Hub
}

func (p *Peers) DispatchKeyFrame() {

}
