'use client'

import Link from 'next/link'
import { usePathname } from 'next/navigation'

const links = [
  { href: '/', label: 'Painel' },
  { href: '/criar_entrada', label: '+ Entrada' },
  { href: '/historico', label: 'Histórico' },
]

export default function Navbar() {
  const pathname = usePathname()

  return (
    <header className="sticky top-0 z-40 border-b ">
      <nav className="mx-auto flex h-14 max-w-5xl items-center gap-1 px-4 sm:gap-2 sm:px-6">
        <span className="mr-auto pr-4 text-sm font-semibold text-foreground sm:text-base">
          Fuel Terminal
        </span>
        {links.map(({ href, label }) => {
          const active = pathname === href
          return (
            <Link
              key={href}
              href={href}
              className={`rounded-lg px-3 py-2 text-sm font-medium ${
                active
                  ? 'shadow-sm'
                  : ''
              }`}
            >
              {label}
            </Link>
          )
        })}
      </nav>
    </header>
  )
}
