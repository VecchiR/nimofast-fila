export type Status = "AGUARDANDO" | "CARREGANDO" | "FINALIZADO" | "CANCELADO"

export type Motorista = {
  id: number
  nome: string
  cpf: string
  cnh: string
  placa: string
}

export type Produto = {
  id: number
  nome: string
}

export type EntradaFila = {
  id: number
  motorista_id: number
  motorista: Motorista
  produto_id: number
  produto: Produto
  status: Status
  horario_chegada: string
  inicio_carregamento?: string
  fim_carregamento?: string
}

export type CriarEntradaFormData = {
  nome: string
  cpf: string
  cnh: string
  placa: string
  produto_id: number
}