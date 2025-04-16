import axios from "axios";
import { AnalysisResult } from "../types/analysis";

const BASE_URL = "http://localhost:8080";

export const analyzePage = async (url: string): Promise<AnalysisResult> => {
  const res = await axios.get(`${BASE_URL}/analyze`, {
    params: { url },
    headers: {
      Authorization: "ApiKey supersecretkey123", // Match .env key in backend
    },
  });
  return res.data;
};
