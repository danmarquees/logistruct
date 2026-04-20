import React, { useMemo } from 'react';
import { motion } from 'framer-motion';
import { AlertTriangle, CheckCircle, PackageSearch } from 'lucide-react';
import { clsx } from 'clsx';
import { twMerge } from 'tailwind-merge';

// Utilitário para classes CSS dinâmicas
const cn = (...inputs) => twMerge(clsx(inputs));

export function WarehouseMap({ insumos }) {
  // Simulando posições fixas (Corredores e Prateleiras) do Digital Twin
  const mappedInsumos = useMemo(() => {
    return insumos.map((item, idx) => {
      // Definir um 'Estoque Físico Atual' aleatório (MOCK) para fins visuais comparativos com o Safety Stock
      // Vamos simular que o estoque de alguns itens está criticamente baixo para vermos o alerta
      const mockCurrentStock = Math.floor(Math.random() * (item.max_sales * 2));
      const isCritical = mockCurrentStock <= item.safety_stock;
      
      // Corredores A, B, C (3 colunas prateleiras)
      const bays = ['A', 'B', 'C', 'D'];
      const currentBay = bays[idx % bays.length];
      const shelfLevel = Math.floor(idx / bays.length) + 1;

      return {
        ...item,
        currentStock: mockCurrentStock,
        isCritical,
        bay: currentBay,
        shelf: shelfLevel,
      };
    });
  }, [insumos]);

  if (!insumos || insumos.length === 0) {
    return (
      <div className="flex flex-col items-center justify-center h-64 text-gray-500">
        <PackageSearch className="w-12 h-12 mb-3 opacity-50" />
        <p>Nenhum insumo mapeado no galpão.</p>
      </div>
    );
  }

  // Agrupar por corredores
  const bays = ['A', 'B', 'C', 'D'];

  return (
    <motion.div 
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      exit={{ opacity: 0, y: -20 }}
      className="p-8 bg-[#0a0f18] rounded-2xl border border-gray-800 shadow-2xl overflow-x-auto"
    >
      <div className="flex flex-col min-w-max">
        {/* Parede superior Galpão */}
        <div className="h-4 w-full bg-slate-800 mb-12 rounded-full relative">
          <span className="absolute -top-6 left-4 text-xs font-bold text-slate-500 tracking-widest uppercase">
            Portões Principais - Zona de Doca
          </span>
        </div>

        <div className="flex gap-16 px-8">
          {bays.map(bayId => {
            const itemsInBay = mappedInsumos.filter(i => i.bay === bayId);
            if (itemsInBay.length === 0) return null;

            return (
              <div key={bayId} className="flex flex-col items-center">
                <div className="w-full text-center border-b border-gray-700 pb-2 mb-6">
                  <h3 className="text-xl font-black text-gray-400">CORREDOR {bayId}</h3>
                </div>

                <div className="space-y-6">
                  {itemsInBay.map((item) => (
                    <motion.div 
                      key={item.id}
                      whileHover={{ scale: 1.05 }}
                      className={cn(
                        "w-64 h-32 rounded-xl p-4 relative group cursor-pointer transition-colors shadow-lg border-2",
                        item.isCritical 
                          ? "bg-red-900/20 border-red-500 animate-pulse-slow" 
                          : "bg-emerald-900/20 border-emerald-500/50 hover:border-emerald-400"
                      )}
                    >
                      {item.isCritical ? (
                        <div className="absolute -top-3 -right-3 bg-red-500 rounded-full p-1.5 shadow-[0_0_15px_rgba(239,68,68,0.7)] z-10">
                          <AlertTriangle className="w-4 h-4 text-white" />
                        </div>
                      ) : (
                        <div className="absolute -top-3 -right-3 bg-emerald-500 rounded-full p-1.5 z-10">
                          <CheckCircle className="w-4 h-4 text-white" />
                        </div>
                      )}
                      
                      <div className="flex flex-col h-full justify-between">
                        <div>
                          <p className="text-xs font-bold text-gray-400 mb-1">PRATELEIRA {item.shelf}</p>
                          <h4 className="text-white font-bold leading-tight line-clamp-2">{item.nome}</h4>
                        </div>
                        
                        <div className="flex justify-between items-end border-t border-gray-700/50 pt-2 mt-2">
                          <div className="flex flex-col">
                            <span className="text-[10px] uppercase text-gray-500">Estoque (Mock)</span>
                            <span className={cn(
                              "font-mono font-bold text-lg leading-none",
                              item.isCritical ? "text-red-400" : "text-emerald-400"
                            )}>
                              {item.currentStock}
                            </span>
                          </div>
                          <div className="flex flex-col text-right">
                            <span className="text-[10px] uppercase text-gray-500">Min. Seg.</span>
                            <span className="font-mono text-sm leading-none text-blue-300">
                              {item.safety_stock}
                            </span>
                          </div>
                        </div>
                      </div>
                    </motion.div>
                  ))}
                </div>
              </div>
            );
          })}
        </div>

      </div>
    </motion.div>
  );
}
