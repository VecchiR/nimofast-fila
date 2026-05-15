import { statusLabels, statusStyles } from "@/lib/statusConfig";
import type { Status } from "@/types"

type Props = {
  status: Status
}

export default function StatusBadge({ status }: Props) {
  return (
    <span className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium border ${statusStyles[status]}`}>
      {statusLabels[status]}
    </span>
  )
}
