export type AnalysisResult = {
  htmlVersion: string;
  title: string;
  headings: Record<string, number>;
  internalLinks: number;
  externalLinks: number;
  inaccessibleLinks: number;
  hasLoginForm: boolean;
};
