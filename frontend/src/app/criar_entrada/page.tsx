"use client"

import { useEffect, useState } from "react"
import { useRouter } from "next/navigation"
import { listarProdutos, criarEntrada } from "@/lib/api"
import { Produto } from "@/types"

export default function CriarEntradaPage() {
  const router = useRouter()

  const [nome,      setNome]      = useState("")
  const [cpf,       setCpf]       = useState("")
  const [cnh,       setCnh]       = useState("")
  const [placa,     setPlaca]     = useState("")
  const [produtoId, setProdutoId] = useState(0)
  const [produtos,  setProdutos]  = useState<Produto[]>([])
  const [erro,      setErro]      = useState("")

  useEffect(() => {
    listarProdutos().then(setProdutos)
  }, [])

  async function handleSubmit(e: React.SubmitEvent) {
    e.preventDefault()
    setErro("")
    try {
      await criarEntrada({ nome, cpf, cnh, placa, produto_id: produtoId })
      router.push("/")
    } catch (e) {
      setErro(e instanceof Error ? e.message : "Erro ao criar entrada")
    }
  }

  return (
    <div className="max-w-md mx-auto bg-white p-6 shadow">
      <h1 className="text-xl font-bold mb-4">Entrada na Fila</h1>

      <form onSubmit={handleSubmit} className="space-y-3">
        <input required value={nome}  onChange={e => setNome(e.target.value)}
          placeholder="Nome" className="w-full border p-2" />
        <input required value={cpf}   onChange={e => setCpf(e.target.value)}
          placeholder="CPF" className="w-full border p-2" />
        <input required value={cnh}   onChange={e => setCnh(e.target.value)}
          placeholder="CNH" className="w-full border p-2" />
        <input required value={placa} onChange={e => setPlaca(e.target.value.toUpperCase())}
          placeholder="Placa" className="w-full border p-2" />

        <select required value={produtoId} onChange={e => setProdutoId(Number(e.target.value))}
          className="w-full border p-2  bg-white">
          <option value={0} disabled>Selecione o produto...</option>
          {produtos.map(p => (
            <option key={p.id} value={p.id}>{p.nome}</option>
          ))}
        </select>

        {erro && <p className="text-red-500 text-sm">{erro}</p>}

        <button type="submit" className="w-full bg-gray-300 p-2 ">
          Criar Entrada
        </button>
      </form>
    </div>
  )
}
