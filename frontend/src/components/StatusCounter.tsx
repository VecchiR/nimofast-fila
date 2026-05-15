import { EntradaFila, Status } from "@/types";
import { statusLabels, statusStyles } from "@/lib/statusConfig";

type Props = {
    fila: EntradaFila[]
    statusToCount: Status[]
}

export default function StatusCounters({fila, statusToCount}: Props) {
  function contarMotoristasPorStatus(entradas: EntradaFila[], status: Status): number {
    return entradas.filter((e) => e.status === status).length;
  }

  return (
    <div className="grid grid-cols-2 gap-3 mb-6 sm:grid-cols-4">
      {(statusToCount as Status[]).map((status) => {
        const count = contarMotoristasPorStatus(fila, status);
        return (
          <div key={status} className={`rounded-lg border p-4 ${statusStyles[status]}`}>
            <p className="text-2xl font-bold">{count}</p>
            <p className="text-sm text-gray-600">
              {statusLabels[status]}
            </p>
          </div>
        );
      })}
    </div>
  );
}
