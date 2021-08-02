### Em quais casos utilizar
- Ideal para microserviços
- Mobile, brownsers e backend
- Geração das bibliotecas de forma automática (stubs)
- Streaming bidirecional utilizando HTTP/2

### Linguagens (Suporte Oficial)
- gRPC-GO
- gRPC-JAVA
- gRPC-C
  - C++
  - Pyhon
  - Ruby
  - Objective C
  - PHP
  - C#
  - Node.js
  - Dart
  - Kotlin / JVM

### Protocol Buffers
- Protocol buffers are Google's language-neutral, platform-neutral, extensible mechanim for serializin structured data - think XML, but smaller, faster, and simpler.

https://developers.google.com/protocol-buffers

##### Protocol Buffers vs JSON
- Arquivos binários < JSON
- Processo de serialização é mais leve (CPU) do que JSON
- Gasta menos recurso de rede
- Processo é mais veloz

Example.proto
````
syntax  = "proto3";

message SearchRequest {
    string query = 1;
    int32 page_number = 2;
    int32 result_per_page = 3;
}
````

#### HTTP/2

- Nome original criado pela Google era SPDY
- Lançado em 2015
- Dados trafegados são binários e não texto como no HTTP 1.1
- Utiliza a mesma conexão TCP para enviar e receber dados do cliente e do servidor (Multiplex)
- Server Push
- Headers são comprimidos
- Gasta menos recursos de rede
- Processo é mais veloz

#### REST vs gRPC

- REST
  - Texto / Json
  - Unidirecional
  - Alta latência
  - Sem contrato
  - Sem suporte a streaming
  - Design pré-definido
  - Bibliotecas de terceiro
- gRPC
  - Protocol Buffers
  - Bidirecional e Assíncrono
  - Baixa latência
  - Contrato definido (.proto)
  - Suporte a straming
  - Design é livre
  - Geração de código