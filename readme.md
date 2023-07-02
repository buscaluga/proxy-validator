# Proxy Validator

O `Proxy Validator` é uma aplicação que tem como objetivo validar uma lista de proxies, filtrar, ordenar por latência e exibir o resultado do seu novo IP. Tem suporte para socks4, socks5, http e https.

## Rodando localmente

``` shell
# Insira sua lista no arquivo de proxies main.go
go run main.go
```

## Exemplo de Resultado
``` shell
Valid Proxies sorted by latency:
Proxy                           Latency         Status   Result
http://95.56.254.139:3128       1.636054437s    true     95.56.254.139
http://35.213.91.45:80          1.85427583s     true     35.194.104.49
http://131.153.48.254:8080      1.981244278s    true     131.153.48.254
http://88.214.41.251:3128       3.243354214s    true     88.214.41.251
http://201.91.82.155:3128       3.244239283s    true     201.91.82.155
http://41.76.145.136:8080       3.665173425s    true     41.76.145.136
http://41.76.145.136:3128       3.668604817s    true     41.76.145.136
http://38.156.238.63:999        4.239560194s    true     38.156.238.63
http://201.182.251.142:999      4.371230862s    true     201.182.251.142
http://185.82.99.42:9093        4.634444308s    true     185.82.98.73
```