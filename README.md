# test-servers
TCP &amp; UDP Test echo servers for testing clustered IPVS/NAT services

Both servers reply on the incoming request with a timestamped message.

# UDP

Usage:

``` sh
$ udp -port 33033
```

``` sh
$ udp -host 127.0.0.1 -port 33033
```
  
# TCP

Usage:

``` sh
$ tcp -port 33033
```
  
``` sh
$ tcp -host 127.0.0.1 -port 33033
```
