import React, { useState, useEffect } from 'react';
import { Package, Activity, AlertCircle, RefreshCw, Trash2 } from 'lucide-react';

function App() {
  const [insumos, setInsumos] = useState([]);
  const [loading, setLoading] = useState(true);

  const fetchInsumos = async () => {
    try {
      setLoading(true);
      const res = await fetch('http://localhost:8080/api/insumos');
      if (!res.ok) throw new Error('Falha ao buscar insumos');
      const data = await res.json();
      setInsumos(data || []);
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchInsumos();
  }, []);

  const deleteInsumo = async (id) => {
    try {
      await fetch(`http://localhost:8080/api/insumos/${id}`, { method: 'DELETE' });
      fetchInsumos();
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-gray-800 to-black p-6">
      <header className="max-w-6xl mx-auto flex items-center justify-between mb-10">
        <div className="flex items-center gap-3">
          <div className="bg-blue-500/20 p-2 rounded-lg">
            <Package className="text-blue-400 w-8 h-8" />
          </div>
          <h1 className="text-3xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-blue-400 to-emerald-400">
            LogiStruct
          </h1>
        </div>
        <button 
          onClick={fetchInsumos}
          className="flex items-center gap-2 px-4 py-2 bg-gray-800 hover:bg-gray-700 border border-gray-700 rounded-lg transition-colors"
        >
          <RefreshCw className={`w-5 h-5 ${loading ? 'animate-spin' : ''}`} />
          Atualizar
        </button>
      </header>

      <main className="max-w-6xl mx-auto">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {insumos.map((item) => (
            <div key={item.id} className="glass-panel p-6 relative group transition-transform hover:-translate-y-1">
              <button 
                onClick={() => deleteInsumo(item.id)}
                className="absolute top-4 right-4 text-gray-500 hover:text-red-400 opacity-0 group-hover:opacity-100 transition-opacity"
              >
                <Trash2 className="w-5 h-5" />
              </button>
              
              <h3 className="text-xl font-semibold mb-4 text-gray-100">{item.nome}</h3>
              
              <div className="space-y-3">
                <div className="flex justify-between text-sm">
                  <span className="text-gray-400">Venda Máxima</span>
                  <span className="font-medium">{item.max_sales} un</span>
                </div>
                <div className="flex justify-between text-sm">
                  <span className="text-gray-400">Lead Time Máx</span>
                  <span className="font-medium">{item.max_lead_time} dias</span>
                </div>
                
                <div className="h-px bg-gray-700 my-4"></div>
                
                <div className="flex items-center justify-between bg-blue-500/10 rounded-lg p-3 border border-blue-500/20">
                  <div className="flex items-center gap-2">
                    <Activity className="text-blue-400 w-5 h-5" />
                    <span className="text-blue-100 font-medium">Estoque de Segurança</span>
                  </div>
                  <span className="text-2xl font-bold text-blue-400">{item.safety_stock}</span>
                </div>
              </div>
            </div>
          ))}
          
          {insumos.length === 0 && !loading && (
            <div className="col-span-full text-center py-12 text-gray-500 flex flex-col items-center">
              <AlertCircle className="w-12 h-12 mb-4 opacity-50" />
              <p>Nenhum insumo encontrado.</p>
            </div>
          )}
        </div>
      </main>
    </div>
  );
}

export default App;
