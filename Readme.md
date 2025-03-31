# DevBook

DevBook é um sistema de postagem desenvolvido para compartilhar ideias, projetos e interações entre desenvolvedores.

## Tecnologias Utilizadas

- **Backend**: Go (Golang) + GORM
- **Banco de Dados**: PostgreSQL
- **Frontend**: HTML, CSS, JavaScript (Framework a definir)
- **Autenticação**: JWT
- **Conversão de HTML para PDF**: Gotenberg v6

## Funcionalidades

- Criação e edição de postagens
- Curtidas e comentários em postagens
- Autenticação de usuários via JWT
- Listagem de postagens
- Exportação de conteúdo em PDF

## Instalação e Execução

### Requisitos
- Go instalado
- PostgreSQL configurado
- Docker (opcional para ambiente de desenvolvimento)

### Passos
1. Clone o repositório:
   ```sh
   git clone https://github.com/seu-usuario/devbook.git
   cd devbook
   ```
2. Configure o banco de dados no arquivo `.env`
3. Execute as migrações do banco:
   ```sh
   go run main.go migrate
   ```
4. Inicie a aplicação:
   ```sh
   go run main.go
   ```

## Contribuição

Fique à vontade para contribuir! Envie um pull request ou abra uma issue para sugerir melhorias.