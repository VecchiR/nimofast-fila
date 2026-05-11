const BASE_URL = 'http://localhost:8080/api/v1';

export async function listarProdutos() {
  const test = await fetch(`${BASE_URL}/produtos`);

  const data = await test.json();

  console.log('data', data);

  return data;
}

export async function listarEntradas() {
  const test = await fetch(`${BASE_URL}/fila`);

  const data = await test.json();

  console.log('data', data);

  return data;
}
