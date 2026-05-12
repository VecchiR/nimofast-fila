# Nimofast - Sistema de Fila de Carregamento de Combustível

## Diário do projeto
Optei por disponibilizar minhas notas que utilizo conforme faço um projeto em que tenho que estudar coisas novas. O "diário do projeto" está em "diario_de_projeto.pdf"


## Instruções para rodar o projeto
### Pré-requisitos
- Docker
- Docker Compose
- Node.js 18+
- npm

### Iniciando o projeto (banco + backend + frontend)

Após clonar o repositório, inicie um terminal **na raíz do projeto** e execute:

```bash
# para iniciar todos os serviços
npm run fullstack-up

# para encerrar serviços do docker (backend + banco)
npm run stop
```

O comando `npm run fullstack-up` é equivalente a:
```bash
# sobe o banco de dados e o backend
docker compose up -d

# instala as dependências do frontend e o inicia
cd frontend/
npm install
npm run dev
```

### Endereços de Acesso

| Serviço | URL |
| :--- | :--- |
| **Frontend (Next.js)** | [http://localhost:3000](http://localhost:3000) |
| **Backend (Go Fiber)** | [http://localhost:8080](http://localhost:8080) |
| **Banco de Dados** | `localhost:5433` |

---
## Arquitetura

### Backend (Go + Fiber)
- Framework: Fiber v3
- Porta: 8080
- Localização: `/backend`

### Frontend (Next.js)
- Framework: Next.js 16 (React 19)
- Porta: 3000
- Localização: `/frontend`

### Banco de dados (PostgreSQL)
- Imagem: postgres:17-alpine
- Porta: 5433
- Database: fuel_terminal
- Localização: Docker container `fuel_terminal_postgres`

---

## Configuração

### Variáveis de Ambiente

Veja `.env.example` para todas as variáveis disponíveis:

```env
# Database
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=fuel_terminal

# Backend
PORT=8080
FRONTEND_URL=http://localhost:3000

# Frontend
NEXT_PUBLIC_API_URL=http://localhost:8080
```
---
## Decisões sobre regras de negócio
- ### Transições de status permitidas
    O ciclo de vida das entradas na fila foi definido assim:
    ```
    AGUARDANDO ──► CARREGANDO ──► FINALIZADO                   
        └───────────────┴──────► CANCELADO
    ```
    **FINALIZADO** e **CANCELADO** são estados terminais: **não podem ser alterados.**
    - **FINALIZADO** representa que o caminhão **foi carregado e saiu da base**. Se o motorista retorna à base ou vai iniciar um novo procedimento, isso deve ser registrado como uma nova visita e deve gerar uma nova entrada no banco para **preservar o histórico** de entradas.
    - **CANCELADO** representa que o caminhão **saiu da base sem concluir o carregamento**, por qualquer motivo. Segue a mesma lógica: **reverter um cancelamento afetaria o histórico**. Se o motorista quiser tentar novamente, cria-se uma nova entrada.

    Sobre **AGUARDANDO** e **CARREGANDO**:
    - A transição **CARREGANDO → AGUARDANDO** foi intencionalmente **não permitida**. Permitir que um caminhão volte de **CARREGANDO** para **AGUARDANDO** afetaria os timestamps de início e fim do carregamento na tabela `entradas_fila`. Caso ocorra algum problema durante o carregamento, o operador deve optar por manter o status como **CARREGANDO** até a resolução, ou passar para **CANCELADO** se o carregamento não puder ser concluído.
    - Cenários de "**carga parcial**" (ex: caminhão foi carregado pela metade) foram considerados, mas não cabia tratá-los como um estado separado. A decisão de classificar esse cenário como **FINALIZADO** ou **CANCELADO** ficaria a critério do operacional da empresa.

- ### Validação de transições no backend
    O backend **valida se a transição de status solicitada é permitida** antes de executar qualquer atualização no banco. Requisições para `PATCH /fila/:id/status` com uma transição inválida são rejeitadas (ex: tentar mover de **FINALIZADO** para **AGUARDANDO**, ou de **CANCELADO** para **CARREGANDO**).

    A decisão de implementar a validação no backend (e não apenas no frontend) foi feita porque o frontend pode ser contornado e, além disso, outros clientes podem vir a interagir com a API no futuro, então **não é seguro contar com a validação apenas por parte do cliente**.
    
- ### Motorista com duas entradas ativas simultâneas
    Regra implementada conforme especificado: o backend rejeita tentativas de criar uma nova entrada para um motorista que já possua uma entrada com status **AGUARDANDO** ou **CARREGANDO**. A mensagem de erro indica que não foi possível criar a entrada devido ao motorista já possuir uma entrada ativa na fila.
    
    Uma consequência dessa regra é que múltiplos carregamentos de produtos diferentes pelo mesmo motorista na mesma visita são tratados como visitas diferentes: o motorista finaliza o primeiro carregamento e então entra novamente na fila para o segundo. Isso simplifica a modelagem ("1 entrada, 1 produto") e está alinhado com a estrutura sugerida no briefing do projeto.

- ### Atualização automática do painel
    *Este ponto está em desenvolvimento*. 
    
    A próxima iteração prevê a implementação de polling periódico com intervalo configurável.

- ### Modelagem de motorista e veículo como entidade única 

    O modelo atual associa CPF, CNH e a placa do veículo ao registro do motorista, assumindo que cada motorista opera sempre o mesmo veículo e um mesmo veículo nunca é utilizado por mais de um motorista.
    
    Essa é uma simplificação consciente feita de acordo com o briefing do projeto. Um modelo mais fiel ao mundo real separaria as entidades MOTORISTA e VEÍCULO, permitindo que um motorista vincule diferentes placas ao longo do tempo e que o id do veículo seja registrado em cada entrada da fila, preservando o histórico de qual caminhão realizou cada carregamento, mesmo que o motorista troque de veículo.



---
*Projeto feito por Rafael Vecchi Silva para o processo seletivo da Nimofast*

Whatsapp: [(16) 99640-9380 ](https://wa.me/5516996409380)

Linkedin: [rafaelvecchisilva](https://www.linkedin.com/in/rafaelvecchisilva/)

Email: vecchi3108@gmail.com