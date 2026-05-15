'use client';

import { format, parseISO, intervalToDuration } from 'date-fns';
import { ptBR } from 'date-fns/locale';
import { useEffect, useState } from 'react';
import { atualizarStatus, listarEntradas } from '@/lib/api';
import { EntradaFila } from '@/types';
import Link from 'next/link';
import StatusBadge from '../../components/StatusBadge';

export default function PainelFila() {
  const [fila, setFila] = useState<EntradaFila[]>([]);
  const [loading, setLoading] = useState(true);
  const [agora, setAgora] = useState<Date | null>(null);

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
    carregar();
    
    setAgora(new Date());
    const intervalo = setInterval(() => {
      setAgora(new Date());
    }, 60000);

    // Limpa o intervalo se desmontar o component
    return () => clearInterval(intervalo);
  }, []);

  async function handleAvancar(id: number, novoStatus: string) {
    try {
      await atualizarStatus(id, novoStatus);
      await carregar();
    } catch (e) {
      alert(e instanceof Error ? e.message : 'Erro');
    }
  }

  // helper para formatar a data
  const formatDate = (timestamp?: string) => {
    if (!timestamp) return '-';
    try {
      // Converte timestamp para Date e formata como "Dia/Mês/Ano Hora:Minuto"
      return format(parseISO(timestamp), 'dd/MM/yyyy HH:mm', { locale: ptBR });
    } catch (error) {
      return 'Data inválida';
    }
  };

  const calcularTempoEspera = (dataIso?: string) => {
    if (!dataIso || !agora) return '-';

    try {
      const dataChegada = parseISO(dataIso);
      
      // não é para acontecer, mas SE a data de chegada for no futuro, mostra 0min para não quebrar
      if (dataChegada > agora) return '0 min';

      const duracao = intervalToDuration({ start: dataChegada, end: agora });
      
      const dias = duracao.days || 0;
      const horas = (duracao.hours || 0) + (dias * 24);
      const minutos = duracao.minutes || 0;

      // se passar de 60 minutos, mostra "HORA h MINUTO min", senão "MINUTO min"
      if (horas > 0) {
        return `${horas}h ${minutos}min`;
      }
      return `${minutos}min`;
      
    } catch (error) {
      console.error('Erro ao calcular tempo de espera. ' + error);
      return '-';
    }
  };


  if (loading) return <p>Carregando...</p>;

  return (
    <div className="flex flex-col gap-4">
      <h1 className="text-2xl font-bold">Painel</h1>

      {fila.length === 0 ? (
        <p className="text-gray-500">Fila vazia.</p>
      ) : (
        <table className="w-full text-sm">
          <thead className="bg-gray-300">
            <tr>
              <th className="p-3 text-left">Motorista</th>
              <th className="p-3 text-left">Placa</th>
              <th className="p-3 text-left">Produto</th>
              <th className="p-3 text-left">Horário de Chegada</th>
              <th className="p-3 text-left">Tempo de Espera</th>
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
                <td className="p-3">{formatDate(entrada.horario_chegada)}</td>
                <td className="p-3">
                  {calcularTempoEspera(entrada.horario_chegada)}
                </td>
                <td className="p-3">
                  <StatusBadge status={entrada.status}/>
                </td>
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
