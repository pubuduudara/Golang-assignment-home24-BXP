import { useState } from "react";
import "./App.css";
import { analyzePage } from "./services/api";
import { AnalysisResult } from "./types/analysis";
import { URLForm } from "./components/URLForm";
import { ResultCard } from "./components/ResultCard";

function App() {
  const [result, setResult] = useState<AnalysisResult | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleAnalyze = async (url: string) => {
    setLoading(true);
    setError("");
    setResult(null);

    try {
      const data = await analyzePage(url);
      setResult(data);
    } catch (err: any) {
      setError(err.message || "Something went wrong.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container">
      <h1>Web Page Analyzer</h1>
      <URLForm onSubmit={handleAnalyze} loading={loading} />
      {error && (
        <div style={{ color: "red", marginBottom: "12px" }}>{error}</div>
      )}
      {result && <ResultCard result={result} />}
    </div>
  );
}

export default App;
