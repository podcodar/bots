import { config } from "dotenv";

interface Settings {
  DISCORD_TOKEN: string;
  BOT_ID: string;
  DAILY_CHANNEL: string;
  SUPABASE_URL: string;
  SUPABASE_TOKEN: string;
}

// @ts-ignore
export const settings: Settings = config();
