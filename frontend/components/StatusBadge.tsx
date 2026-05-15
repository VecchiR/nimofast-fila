import type { Status } from '@/types'


const styles = {
  AGUARDANDO: 'bg-yellow-100 text-yellow-800 border-yellow-300',
  CARREGANDO: 'bg-blue-100  text-blue-800  border-blue-300',
  FINALIZADO: 'bg-green-100 text-green-800 border-green-300',
  CANCELADO:  'bg-red-100   text-red-800   border-red-300',
}

const labels = {
    AGUARDANDO: '⏳ Aguardando',
    CARREGANDO: '⛽ Carregando',
    FINALIZADO: '✅ Finalizado',
    CANCELADO:  '❌ Cancelado',
}

type Props = {
  status: Status
}

export default function StatusBadge({ status }: Props) {
  return (
    <span className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium border ${styles[status]}`}>
      {labels[status]}
    </span>
  )
}
