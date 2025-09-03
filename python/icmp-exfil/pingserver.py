from scapy.all import *

def callback(pkt):
   if pkt.haslayer(ICMP):
      #print(pkt.summary())
      #print(pkt.show())
      print(pkt[ICMP].load)

rx = sniff(iface="lo", filter="icmp", prn=callback, count=1)
print(rx)
