// NOTICE

// Project Name: Cloaq
// Copyright © 2026 Neil Talap and/or its designated Affiliates.

// This software is licensed under the Dragonfly Public License (DPL) 1.0.

// All rights reserved. The names "Neil Talap" and any associated logos or branding
// are trademarks of the Licensor and may not be used without express written permission,
// except as provided in Section 7 of the License.

// For commercial licensing inquiries or permissions beyond the scope of this
// license, please create an issue in github.

package routing

import (
	"log"

	"cloaq/src/tun"
)

func CreateIPv6PacketListener(dev tun.Device) {
	buf := make([]byte, 65535)
	for {
		n, err := dev.Read(buf)
		if err != nil {
			log.Println("tun.Read error:", err)
			continue
		}

		pkt := buf[:n]
		if len(pkt) < 40 {
			continue
		}
		if (pkt[0] >> 4) != 6 {
			continue
		}

		payload := pkt[40:]
		log.Printf("IPv6 packet: %d bytes, payload %d bytes\n", len(pkt), len(payload))
	}
}
