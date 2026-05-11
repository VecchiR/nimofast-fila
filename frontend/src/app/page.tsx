'use client';

import { useEffect, useState } from 'react';
import { listarEntradas, listarProdutos } from '@/lib/api';

export default function PainelFila() {
  const [produtos, setProdutos] = useState<any[]>([]);
  const [fila, setFila] = useState<any[]>([]);

  useEffect(() => {
    async function fetchProdutos() {
      try {
        const data = await listarProdutos();
        setProdutos(data);
      } catch (error) {
        console.error('Erro ao listar produtos:', error);
      }
    }
    fetchProdutos();
  }, []);

  useEffect(() => {
    async function fetchFila() {
      try {
        const data = await listarEntradas();
        setFila(data);
      } catch (error) {
        console.error('Erro ao buscar a fila:', error);
      }
    }
    fetchFila();
  }, []);

  return (
    <div>
      <h1 className="text-2xl font-bold mb-4">Painel</h1>

      {fila.length === 0 ? (
        <p className="text-gray-500">Fila vazia.</p>
      ) : (
        <table className="w-full text-sm">
          <thead className="bg-gray-300">
            <tr>
              <th className="p-3 text-left">Motorista</th>
              <th className="p-3 text-left">Placa</th>
              <th className="p-3 text-left">Produto</th>
              <th className="p-3 text-left">Status</th>
            </tr>
          </thead>
          <tbody className='bg-gray-200'>
            {fila.map((entrada) => (
              <tr key={entrada.id} className="border-t">
                <td className="p-3">{entrada.motorista.nome}</td>
                <td className="p-3">{entrada.motorista.placa}</td>
                <td className="p-3">{entrada.produto.nome}</td>
                <td className="p-3">{entrada.status}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}
