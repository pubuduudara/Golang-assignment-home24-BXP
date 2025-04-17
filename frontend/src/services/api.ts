import axios from "axios";
import { AnalysisResult, ApiResponse } from "../types/analysis";

const BASE_URL = import.meta.env.VITE_API_URL || "http://localhost:8080";

export const analyzePage = async (url: string): Promise<AnalysisResult> => {
  try {
    const res = await axios.get<ApiResponse>(`${BASE_URL}/analyze`, {
      params: { url },
      headers: {
        "X-API-Key": import.meta.env.VITE_API_KEY,
      },
    });

    const { status, data } = res.data;

    if (!status) {
      throw new Error(data); // backend error message
    }

    return data;
  } catch (error: any) {
    // Extract message if available
    const apiMessage = error?.response?.data?.data;
    throw new Error(apiMessage || error.message || "Unknown error occurred");
  }
};
