# Desafio Técnico Go lang - Caio Milfont

## 🚀 Indice

- 📓 [Sobre](#-Sobre)
- 🧱 [Funcionalidades da aplicação](#-Funcionalidades-da-aplicação)
- 👨‍💻 [Tecnologias utilizadas](#-Tecnologias-utilizadas)
- 📦 [Como baixar o projeto](#-Como-baixar-o-projeto)
- 🤝 [Considerações](#-Considerações)

## 📓 Sobre

Referência do projeto **Desafio Técnico - Go(lang)**

O desafio proposto aqui foi a implementação de uma API que fosse capaz de realizar transferências entre contas internas de um banco digital. 

## 🧱 Funcionalidades da aplicação

### **Contas**

- Obter a lista de Contas presentes no banco digital.

- Obter saldo de uma conta específica.

- Criar uma nova Conta.

### **Login**

- Autenticação de um usuário retornando um token para ser utilizado nas demais requisições da API

### **Transferências**

- Obter a lista de transferências realizadas pelo usuário autenticado (obtido a partir do token gerado no [Login](#-Login))

- Realizar uma transferência entre contas.

## 👨‍💻 Tecnologias utilizadas

### Back End
- [Go Lang](https://go.dev/)

### DevOps
- [Docker](https://www.docker.com/)

---

## 📦 Como baixar o projeto

### Dependências:
- [Docker](https://www.docker.com/)
- [Go Lang](https://go.dev/)

### Comandos:

```bash
 # Clonar o repositório
 git clone https://github.com/mgbcaio/desafio-go-stone.git
```

```bash
 # Entrar no diretório
 cd desafio-go-stone
```

```bash
 # Construir a imagem do Docker e executá-la com o comando make
 make start
```

```bash
 # Ou, construir e executar utilizando os próprios comando do Docker
 docker build -t transfer/app:latest . &&
 docker run -p 9090:9090 transfer/app:latest
```

Executar as chamadas à API no endereço: http://localhost:9090

---

## 🤝 Considerações

Gostaria de agradecer a oportunidade que me foi dada de implementar esse desafio técnico. Confesso que foi bastante produtivo e pude aprender um pouco mais com todas as pesquisas que tive que fazer para implementá-lo. Espero que o código esteja de acordo com o que estão procurando em um desenvolver. No mais, feedback sobre o projeto como um todo são muito bem-vindos. Obrigado! 😁