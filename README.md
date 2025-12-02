# 🤖 Encurtador de URL em Go (Golang)

- Este é um projeto simples de encurtador de URL desenvolvido em Go. Ele permite que os usuários encurtem URLs longas e sejam redirecionados para a URL original usando um shortID exclusivo. O projeto utiliza criptografia AES para armazenar as URLs de forma segura.

## ✨ Funcionalidades

- Encurtamento de URL: Gera um shortID curto e aleatório para qualquer URL longa.

- Redirecionamento: Redireciona o usuário do shortID para a URL original.

- Criptografia Segura: Utiliza o algoritmo AES para criptografar as URLs originais antes de armazená-las, garantindo que o conteúdo sensível não seja armazenado em texto simples.

- Geração Aleatória de ID: O shortID é gerado usando uma combinação alfanumérica aleatória ([a-z][A-Z][0-9]).

- Concorrência Segura: Usa sync.Mutex para gerenciar o acesso concorrente ao armazenamento de URLs (urlStore).

## 🚀 Como Executar o Projeto

Pré-requisitos:
Certifique-se de ter o Go (versão 1.16 ou superior) instalado em sua máquina.

Passos de Execução

Clone o repositório: https://github.com/danrodsg/url-shortener.git

Bash:

- git clone https://github.com/danrodsg/url-shortener.git
- cd url-shortener

Execute o servidor:

go run main.go

O servidor estará rodando em http://localhost:8081

## 🛠️Exemplo de Uso
- Encurtamento (Não implementado na API, apenas na lógica): Uma requisição para encurtar a URL "https://www.google.com" gera um ID, por exemplo, aBc123.

- Redirecionamento: Acessar o link no navegador: http://localhost:8080/r/aBc123

- Resultado Esperado: O navegador será redirecionado para https://www.google.com.



