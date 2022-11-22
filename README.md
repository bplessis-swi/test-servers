# test-servers
TCP &amp; UDP Test echo servers for testing clustered IPVS/NAT services

Both servers reply on the incoming request with a timestamped message.

# UDP

Usage:

  udp -port 33033
  
  udp -host 127.0.0.1 -port 33033
  
  
# TCP

Usage:

  tcp -port 33033
  
  tcp -host 127.0.0.1 -port 33033
