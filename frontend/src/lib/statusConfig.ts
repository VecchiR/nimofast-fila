import { Status } from "@/types";

export const statusStyles: Record<Status, string> = {
  AGUARDANDO: "bg-yellow-100 text-yellow-800 border-yellow-300",
  CARREGANDO: "bg-blue-100  text-blue-800  border-blue-300",
  FINALIZADO: "bg-green-100 text-green-800 border-green-300",
  CANCELADO:  "bg-red-100   text-red-800   border-red-300",
};

export const statusLabels: Record<Status, string> = {
  AGUARDANDO: "⏳ Aguardando",
  CARREGANDO: "⛽ Carregando",
  FINALIZADO: "✅ Finalizado",
  CANCELADO:  "❌ Cancelado",
};