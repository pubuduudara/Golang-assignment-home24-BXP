import axios from "axios";
import { AnalysisResult, ApiResponse } from "../types/analysis";

const BASE_URL = "http://localhost:8080";

export const analyzePage = async (url: string): Promise<AnalysisResult> => {
  const res = await axios.get<ApiResponse>(`${BASE_URL}/analyze`, {
    params: { url },
    headers: {
      "X-API-Key": "supersecretkey123", //TODO: move to env variable
    },
  });

  if (!res.data.status) {
    throw new Error(res.data.data);
  }

  return res.data.data;
};
