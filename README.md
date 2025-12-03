# 🤖 Curtador de URL em Go (Golang)

[![Go](https://github.com/golang/go/blob/master/assets/badge.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Este é um projeto simples e eficiente de **encurtador de URL** desenvolvido em Go. Ele permite que usuários acessem URLs longas complexas sendo redirecionados a partir de um `shortID` exclusivo e curto. O projeto foca em **segurança** e **concorrência**, utilizando criptografia AES para o armazenamento e `sync.Mutex` para acesso seguro aos dados.

## ✨ Funcionalidades Principais

| Recurso | Descrição | Detalhes Técnicos |
| :--- | :--- | :--- |
| **Encurtamento de URL** | Gera um `shortID` curto e aleatório para mapear qualquer URL original. | Combinação alfanumérica de letras e números `[a-z][A-Z][0-9]`. |
| **Redirecionamento Rápido** | Redireciona o usuário do `shortID` para a URL original. | Utiliza o `http.Redirect` do pacote `net/http` em Go. |
| **Criptografia Segura** | As URLs originais são criptografadas antes do armazenamento. | Algoritmo **AES (Advanced Encryption Standard)** é usado para garantir que o conteúdo sensível não seja armazenado em texto simples. |
| **Concorrência Segura** | Gerencia o acesso simultâneo ao armazenamento de URLs. | Utiliza `sync.Mutex` para bloquear e liberar o acesso ao mapa de URLs (`urlStore`), prevenindo *race conditions*. |

---

## 🚀 Como Executar o Projeto

Siga os passos abaixo para configurar e rodar o servidor localmente.

### Pré-requisitos

Certifique-se de ter o **Go** (versão **1.16 ou superior**) instalado em sua máquina.

### Passos de Execução

1.  **Clone o repositório:**
    ```bash
    git clone [https://github.com/danrodsg/url-shortener.git](https://github.com/danrodsg/url-shortener.git)
    cd url-shortener
    ```

2.  **Execute o servidor:**
    ```bash
    go run main.go
    ```

O servidor estará rodando em **`http://localhost:8081`**.

---

## 🛠️ Exemplo de Uso

O projeto foi projetado com dois fluxos principais: o encurtamento (lógica interna) e o redirecionamento (endpoint público).

### 1. Encurtamento (Lógica Interna)

Embora a **API pública de encurtamento não esteja implementada** neste exemplo (ela ocorre apenas na lógica interna/função), o processo interno é o seguinte:

* **URL Original:** `https://www.google.com`
* **Processo:** O sistema criptografa a URL, gera um `shortID` aleatório (ex: `aBc123`), armazena o par (ID -> URL Criptografada) e desbloqueia o acesso.

### 2. Redirecionamento (Endpoint Público)

Para testar o recurso de redirecionamento, simule o acesso a um `shortID` gerado.

* **Endpoint de Redirecionamento:**
    ```
    http://localhost:8081/r/{shortID}
    ```

* **Exemplo de Acesso:** Se o ID gerado for `aBc123`:
    ```
    http://localhost:8081/r/aBc123
    ```

#### Resultado Esperado

Ao acessar o link no navegador, o servidor decifra a URL original (`https://www.google.com`) e o navegador será redirecionado automaticamente para o destino.

---

## 👨‍💻 Estrutura e Tecnologia

O projeto utiliza os seguintes componentes-chave do ecossistema Go:

* **`net/http`:** Para lidar com as rotas do servidor web e o redirecionamento HTTP.
* **`crypto/aes`:** Para a implementação da criptografia de 128/192/256 bits.
* **`sync`:** O `sync.Mutex` é vital para gerenciar o estado compartilhado (`urlStore`) e garantir a **segurança de concorrência** (thread-safety) em um ambiente multi-rotinas (goroutines) típico de aplicações Go.
* **`math/rand`:** Usado na geração pseudo-aleatória do `shortID`.



