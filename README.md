# SonicScan

Super rapid port scanning using Go concurrency


## Setup:
To run build using:
```
go build scanner.go
```
and then run the executable using
```
./scanner
```
## Supported options:
The application currently supports the following flag options:

- ```-n_threads``` This flag can be used to toggle the amount of workers for port scanning. A rule of thumb is to keep this below 200. Default value is 8.
- ```-target``` This flag is used to toggle the target IP address. The default value of is set as 8.8.8.8 (Google's IP address).


## To do:
- [] Use synchronization to avoid closing connection before response
- [] Let user prepare a text file to choose which IP/ports to scan from
- [] Support automatic name server resolver
- [] Exception catching for various scenarios (packets dropped between host & server, connection refused by server, timeout, port is occupied by other process etc)
