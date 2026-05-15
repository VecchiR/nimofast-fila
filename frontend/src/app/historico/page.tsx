"use client";

import { listarHistorico } from "@/lib/api";
import { EntradaFila } from "@/types";
import { useEffect, useState } from "react";
import StatusBadge from "../../../components/StatusBadge";

export default function Historico() {
  const [historico, setHistorico] = useState<EntradaFila[]>([]);
  const [loading, setLoading] = useState(true);

  async function carregar() {
    try {
      setLoading(true);
      const data = await listarHistorico();
      setHistorico(data);
    } catch (e) {
      console.error(e);
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    carregar();
  }, []);

  if (loading) return <p>Carregando...</p>;

  return (
    <div className="flex flex-col gap-4">
      <h1 className="text-2xl font-bold">Histórico</h1>

      {historico.length === 0 ? (
        <p className="text-gray-500">Histórico vazio.</p>
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
          <tbody className="bg-gray-200">
            {historico.map((entrada) => (
              <tr key={entrada.id} className="border-t">
                <td className="p-3">{entrada.motorista.nome}</td>
                <td className="p-3">{entrada.motorista.placa}</td>
                <td className="p-3">{entrada.produto.nome}</td>
                <td className="p-3">
                  <StatusBadge status={entrada.status}/>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}
