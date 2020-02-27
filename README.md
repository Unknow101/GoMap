# GoMap
Small Golang port scanner

The main goal of this project is not to produce a efficient pentest toos ... just to learn Golang Basics !

/!\ Disclaimer : I am not responsable about anything you can do with it !

## Demo

```
gomap.exe -host scanme.nmap.org

 _____       _____
|   __| ___ |     | ___  ___
|  |  || . || | | || .'|| . |
|_____||___||_|_|_||__,||  _|
                        |_|
v0.2 version
==========
Scanning scanme.nmap.org for 1024 ports..
==========
22 Port open
80 Port open
Scan done in 10.2247838s seconds
```

## Documentation

```
Usage of gomap.exe:
  -f    Scan all 65535 ports
  -host string
        Host to scan (default "false")
```

## Exemple

Run full scan :
```gomap.exe -f -host scanme.nmap.org```

Run fast scan :
```gomap.exe -host scanme.nmap.org```
        
     
