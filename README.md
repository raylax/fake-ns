# Fake name server
## Useage
```bash
./fake-ns -d xxx.com
```

## Example
```bash
> nslookup 127.0.0.1.fns.inurl.org 8.8.8.8
Server:		8.8.8.8
Address:	8.8.8.8#53

Non-authoritative answer:
Name:	127.0.0.1.fns.inurl.org
Address: 127.0.0.1

> nslookup 0xffffffff.fns.inurl.org 8.8.8.8
Server:		8.8.8.8
Address:	8.8.8.8#53

Non-authoritative answer:
Name:	0xffffffff.fns.inurl.org
Address: 255.255.255.255

> nslookup 2130706433.fns.inurl.org 8.8.8.8
Server:		8.8.8.8
Address:	8.8.8.8#53

Non-authoritative answer:
Name:	2130706433.fns.inurl.org
Address: 127.0.0.1
```
[![asciicast](https://asciinema.org/a/445514.svg)](https://asciinema.org/a/445514)
