'use client';

import { useEffect, useState } from 'react';
import { atualizarStatus, listarEntradas } from '@/lib/api';
import { EntradaFila } from '@/types';
import Link from 'next/link';

export default function PainelFila() {
  const [fila, setFila] = useState<EntradaFila[]>([]);
  const [loading, setLoading] = useState(true);

  async function carregar() {
    try {
      setLoading(true);
      const data = await listarEntradas();
      setFila(data);
    } catch (e) {
      console.error(e);
    } finally {
      setLoading(false);
    }
  }

useEffect(() => {
  carregar()
}, [])

  async function handleAvancar(id: number, novoStatus: string) {
    try {
      await atualizarStatus(id, novoStatus);
      await carregar();
    } catch (e) {
      alert(e instanceof Error ? e.message : 'Erro');
    }
  }

  if (loading) return <p>Carregando...</p>;

  return (
    <div className="flex flex-col gap-4">
      <h1 className="text-2xl font-bold">Painel</h1>

      <Link href="/historico" className="self-start bg-blue-300 p-2">
        Histórico
      </Link>

      <Link href="/criar_entrada" className="self-start bg-blue-300 p-2">
        + Nova Entrada
      </Link>

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
              <th className="p-3 text-left">Ações</th>
            </tr>
          </thead>
          <tbody className="bg-gray-200">
            {fila.map((entrada) => (
              <tr key={entrada.id} className="border-t">
                <td className="p-3">{entrada.motorista.nome}</td>
                <td className="p-3">{entrada.motorista.placa}</td>
                <td className="p-3">{entrada.produto.nome}</td>
                <td className="p-3">{entrada.status}</td>
                <td className="p-3 flex gap-2">
                  {entrada.status === 'AGUARDANDO' && (
                    <button
                      onClick={() => handleAvancar(entrada.id, 'CARREGANDO')}
                      className="bg-blue-500 text-white p-2 text-xs"
                    >
                      Iniciar
                    </button>
                  )}
                  {entrada.status === 'CARREGANDO' && (
                    <button
                      onClick={() => handleAvancar(entrada.id, 'FINALIZADO')}
                      className="bg-green-500 text-white p-2 text-xs"
                    >
                      Finalizar
                    </button>
                  )}
                  {(entrada.status === 'AGUARDANDO' || entrada.status === 'CARREGANDO') && (
                    <button
                      onClick={() => handleAvancar(entrada.id, 'CANCELADO')}
                      className="bg-red-400 text-white p-2 text-xs"
                    >
                      Cancelar
                    </button>
                  )}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}
