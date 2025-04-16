import { AnalysisResult } from "../types/analysis";

export const ResultCard = ({ result }: { result: AnalysisResult }) => (
  <div className="result">
    <div>
      <strong>HTML Version:</strong> {result.htmlVersion}
    </div>
    <div>
      <strong>Title:</strong> {result.title}
    </div>
    <div>
      <strong>Headings:</strong>
      <ul>
        {Object.entries(result.headings).map(([h, count]) => (
          <li key={h}>
            {h.toUpperCase()}: {count}
          </li>
        ))}
      </ul>
    </div>
    <div>
      <strong>Internal Links:</strong> {result.internalLinks}
    </div>
    <div>
      <strong>External Links:</strong> {result.externalLinks}
    </div>
    <div>
      <strong>Inaccessible Links:</strong> {result.inaccessibleLinks}
    </div>
    <div>
      <strong>Contains Login Form:</strong> {result.hasLoginForm ? "Yes" : "No"}
    </div>
  </div>
);
