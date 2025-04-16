export type AnalysisResult = {
  htmlVersion: string;
  title: string;
  headings: Record<string, number>;
  links: {
    internal: number;
    external: number;
    inaccessible: number;
  };
  hasLoginForm: boolean;
};

export type ApiResponse =
  | { status: true; data: AnalysisResult }
  | { status: false; data: string };
