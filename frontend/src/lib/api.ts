import type { Motorista, Produto, EntradaFila, CriarEntradaFormData, Status } from "@/types";


const BASE_URL = 'http://localhost:8080/api/v1';

// concentra as chamadas em um helper para deixar mais limpo e tratar de erros em um lugar só
async function fetcher<T>(endpoint: string, options?: RequestInit): Promise<T> {
  const response = await fetch(`${BASE_URL}${endpoint}`, {
    headers: {
      'Content-Type': 'application/json',
    },
    ...options,
  });

  const jsonData = await response.json();
  
  if (!response.ok) {
    // o backend sempre manda {error: ....}, então tento usar isso, mas se não der -> mensagem genérica de erro http
    throw new Error(jsonData.error ?? `HTTP ERROR - status: ${response.status}`);
  }


  return jsonData as T;
}

export async function listarProdutos(): Promise<Produto[]>{
  return fetcher('/produtos')
}

export async function listarEntradas(): Promise<EntradaFila[]> {
  return fetcher(`/fila`);
}

export function criarEntrada(formData: CriarEntradaFormData) {
  return fetcher('/fila', {
    method: 'POST',
    body: JSON.stringify(formData),
  })
}