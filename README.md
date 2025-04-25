# API de Vagas de Trabalho

## Descrição

GOstd é uma API desenvolvida em Go para gerenciar vagas de trabalho. O projeto foi criado para facilitar o cadastro, consulta e gerenciamento de vagas, permitindo que empresas publiquem oportunidades e candidatos encontrem posições relevantes.

## Funcionalidades

- **Cadastro de Vagas:** Permite que empresas criem e editem vagas.
- **Consulta de Vagas:** Listagem e busca de vagas disponíveis.
- **Filtros:** Pesquisa por palavras-chave, localização, e tipo de trabalho.
- **API RESTful:** Endpoints bem definidos para integração com outros sistemas.

## Como Rodar o Projeto

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/thiagoeu/GOstd.git
   ```

2. **Entre no diretório:**

   ```bash
   cd GOstd
   ```

3. **Instale as dependências:**

   ```bash
   go mod tidy
   ```

4. **Execute a aplicação:**

   ```bash
   go run main.go
   ```

5. **Acesse a API:**

   A aplicação estará rodando em:

   ```
   http://localhost:8080
   ```

## Endpoints

- `GET /jobs` — Lista todas as vagas.
- `POST /jobs` — Cria uma nova vaga.
- `GET /jobs/{id}` — Detalhes de uma vaga específica.
- `PUT /jobs/{id}` — Atualiza uma vaga existente.
- `DELETE /jobs/{id}` — Remove uma vaga.

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença

Este projeto está sob a licença MIT. Consulte o arquivo LICENSE para mais detalhes.
