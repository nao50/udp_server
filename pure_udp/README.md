# UDP study

## Setup
```
$ git clone https://github.com/naoyamaguchi/udp_server.git
$ cd udp_server
$ docker build ./ -t udpserver:latest
$ docker run -it -d --security-opt=apparmor=unconfined --cap-add=SYS_PTRACE --cap-add=NET_ADMIN  --cap-add=NET_RAW udpserver:latest
$ docker images
$ docker exec -it <CONTAINER ID> bash
$ cd pure_udp/
$ go build udpserver.go && strace -e 'trace=!pselect6,futex,sched_yield' ./udpserver
```

output
```
--- snip ---
epoll_create1(EPOLL_CLOEXEC)            = 4
epoll_ctl(4, EPOLL_CTL_ADD, 3, {EPOLLIN|EPOLLOUT|EPOLLRDHUP|EPOLLET, {u32=1772719872, u64=140515922775808}}) = 0
fcntl(3, F_GETFL)                       = 0x8000 (flags O_RDONLY|O_LARGEFILE)
fcntl(3, F_SETFL, O_RDONLY|O_NONBLOCK|O_LARGEFILE) = 0
read(3, "128\n", 65536)                 = 4
read(3, "", 65532)                      = 0
epoll_ctl(4, EPOLL_CTL_DEL, 3, 0xc42004fd4c) = 0
close(3)  

===== net.ListenUDP() =====

socket(AF_INET, SOCK_DGRAM|SOCK_CLOEXEC|SOCK_NONBLOCK, IPPROTO_IP) = 3
setsockopt(3, SOL_SOCKET, SO_BROADCAST, [1], 4) = 0
bind(3, {sa_family=AF_INET, sin_port=htons(2152), sin_addr=inet_addr("127.0.0.1")}, 16) = 0
epoll_ctl(4, EPOLL_CTL_ADD, 3, {EPOLLIN|EPOLLOUT|EPOLLRDHUP|EPOLLET, {u32=1772719872, u64=140515922775808}}) = 0
getsockname(3, {sa_family=AF_INET, sin_port=htons(2152), sin_addr=inet_addr("127.0.0.1")}, [112->16]) = 0

===== updConn.ReadFromUDP() =====

recvfrom(3, 0xc42004f8ba, 1550, 0, 0xc42004f5d8, [112]) = -1 EAGAIN (Resource temporarily unavailable)
epoll_pwait(4, [{EPOLLOUT, {u32=4207877888, u64=140097451138816}}], 128, 0, NULL, 140097451138816) = 1
epoll_pwait(4, 

```
At same container
```
$ go run udpclient.go
```

Move on
```
epoll_pwait(4, [{EPOLLIN|EPOLLOUT, {u32=4207877888, u64=140097451138816}}], 128, -1, NULL, 140097451138816) = 1
recvfrom(3, "0123456789", 1550, 0, {sa_family=AF_INET, sin_port=htons(41343), sin_addr=inet_addr("127.0.0.1")}, [112->16]) = 10
write(1, "recieved size:  10\n", 19recieved size:  10
)    = 19
write(1, "remote address:  127.0.0.1:41343"..., 33remote address:  127.0.0.1:41343
) = 33
write(1, "\n===== updConn.ReadFromUDP() ==="..., 35
===== updConn.ReadFromUDP() =====
) = 35
recvfrom(3, 0xc42004f8ba, 1550, 0, 0xc42004f5d8, [112]) = -1 EAGAIN (Resource temporarily unavailable)
epoll_pwait(4, [], 128, 0, NULL, 140097451138816) = 0
epoll_pwait(4, recvfrom(3, 0xc42004f8ba, 1550, 0, 0xc42004f5d8, [112]) = -1 EAGAIN (Resource temporarily unavailable)
epoll_pwait(4, [{EPOLLOUT, {u32=4207877888, u64=140097451138816}}], 128, 0, NULL, 140097451138816) = 1
epoll_pwait(4,

```


## Study
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


```
$ go build udpclient.go && strace -e 'trace=!pselect6,futex,sched_yield' ./udpclient
```

```
openat(AT_FDCWD, "/proc/sys/net/core/somaxconn", O_RDONLY|O_CLOEXEC) = 3
epoll_create1(EPOLL_CLOEXEC)            = 4
epoll_ctl(4, EPOLL_CTL_ADD, 3, {EPOLLIN|EPOLLOUT|EPOLLRDHUP|EPOLLET, {u32=3262414592, u64=139821627768576}}) = 0
fcntl(3, F_GETFL)                       = 0x8000 (flags O_RDONLY|O_LARGEFILE)
fcntl(3, F_SETFL, O_RDONLY|O_NONBLOCK|O_LARGEFILE) = 0
read(3, "128\n", 65536)                 = 4
read(3, "", 65532)                      = 0
epoll_ctl(4, EPOLL_CTL_DEL, 3, 0xc42004fd4c) = 0
close(3)                                = 0
write(1, "\n===== net.Dial() =====\n", 24
===== net.Dial() =====
) = 24
socket(AF_INET, SOCK_DGRAM|SOCK_CLOEXEC|SOCK_NONBLOCK, IPPROTO_IP) = 3
setsockopt(3, SOL_SOCKET, SO_BROADCAST, [1], 4) = 0
connect(3, {sa_family=AF_INET, sin_port=htons(2152), sin_addr=inet_addr("127.0.0.1")}, 16) = 0
epoll_ctl(4, EPOLL_CTL_ADD, 3, {EPOLLIN|EPOLLOUT|EPOLLRDHUP|EPOLLET, {u32=3262414592, u64=139821627768576}}) = 0
getsockname(3, {sa_family=AF_INET, sin_port=htons(36044), sin_addr=inet_addr("127.0.0.1")}, [112->16]) = 0
getpeername(3, {sa_family=AF_INET, sin_port=htons(2152), sin_addr=inet_addr("127.0.0.1")}, [112->16]) = 0
write(1, "\n===== conn.Write() =====\n", 26
===== conn.Write() =====
) = 26
write(3, "0123456789", 10)              = 10
epoll_ctl(4, EPOLL_CTL_DEL, 3, 0xc42004fd94) = 0
close(3)                                = 0
exit_group(0)                           = ?
+++ exited with 0 +++
```

```
$ go build rawserver.go && strace -e 'trace=!pselect6,futex,sched_yield' ./rawserver
```
