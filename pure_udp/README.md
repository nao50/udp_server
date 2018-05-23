# UDP

```
func ListenPacket(network, address string) (PacketConn, error)
```

```
func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)

type UDPAddr struct {
        IP   IP
        Port int
        Zone string // IPv6 scoped addressing zone
}

https://golang.org/pkg/net/#UDPConn

```