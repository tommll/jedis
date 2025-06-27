Quesion: what is RESP?
- RESP is a protocol Redis use to communicate between client and server


For example, to communicate string data
```
$5\r\nhello\r\n       -- “hello”
$0\r\n\r\n            -- empty string
$-1\r\n               -- null bulk string
```

integer data
```
:1000\r\n
```
