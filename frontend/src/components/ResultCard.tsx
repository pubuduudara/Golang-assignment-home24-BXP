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
        {Object.entries(result.headings).map(([tag, count]) => (
          <li key={tag}>
            {tag.toUpperCase()}: {count}
          </li>
        ))}
      </ul>
    </div>
    <div>
      <strong>Links:</strong>
      <ul>
        <li>Internal: {result.links.internal}</li>
        <li>External: {result.links.external}</li>
        <li>Inaccessible: {result.links.inaccessible}</li>
      </ul>
    </div>
    <div>
      <strong>Contains Login Form:</strong> {result.hasLoginForm ? "Yes" : "No"}
    </div>
  </div>
);
