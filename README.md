# Yamaha RX-V7xx YCAST Server
Since 2019 the official Yamaha Webradio Browser (http://radioyamaha.vtuner.com) is charing money for their service.
However, this has not lead to a better and faster service. 
Thus, i decided to move away from them and implement a simple server of my own.

All you need to do is to bogus the DNS of the official vtuner IPs (radioyamaha.vtuner.com and radioyamaha2.vtuner.com) to your own IPs.
The simplest way may be to do this in your router (not many will support this).
Another way is to run an own DNS server. "Pi-Hole" may be a good for a first start, i use it myself due to its simplicity!

Last step is to setup the DNS server IP in your RX-V Receiver to the one of your own DNS server.

