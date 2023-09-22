package main

import (
	"encoding/json"
	"github.com/omidfth/godp"
	"log"
)

func main() {
	r := godp.NewRouter()
	r.NewRoute(godp.PING, OnPing)
	r.ListenAndServe(1055, godp.MaxBufferSize)
}

func OnPing(c *godp.Context) {
	log.Println("On Ping")
	c.GetPingService().Ping(c.Packet.SocketID)
	packet := godp.MakePacket(c.Packet.SocketID, c.Packet.RoomID, godp.PING, nil)
	j, _ := json.Marshal(packet)
	c.Emit(c.Address, j)
}
